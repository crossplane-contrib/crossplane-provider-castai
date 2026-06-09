#!/usr/bin/env python3
"""Transform nested_type attributes in schema.json to block_types.

Upjet < v2 cannot handle the Terraform Plugin Framework nested_type attribute
format. This script converts nested_type attrs to the older block_types format
so that upjet v1.x can generate types correctly.
"""

import json
import sys


def transform_block(block):
    attrs = block.get("attributes", {})
    block_types = block.get("block_types", {})
    new_attrs = {}
    new_block_types = dict(block_types)

    for name, attr in attrs.items():
        if "nested_type" in attr:
            nt = attr["nested_type"]
            nesting_mode = nt.get("nesting_mode", "single")
            inner_block = {"attributes": nt.get("attributes", {})}
            transform_block(inner_block)
            inner_block_result = {
                    "attributes": inner_block["attributes"],
                    "description_kind": "plain",
            }
            if "block_types" in inner_block:
                inner_block_result["block_types"] = inner_block["block_types"]
            converted = {
                "block": inner_block_result,
                "nesting_mode": nesting_mode,
                "min_items": 0,
                "max_items": 1 if nesting_mode == "single" else 0,
            }
            if attr.get("description"):
                converted["block"]["description"] = attr["description"]
            new_block_types[name] = converted
        else:
            new_attrs[name] = attr

    block["attributes"] = new_attrs
    if new_block_types:
        block["block_types"] = new_block_types


def relax_conditionally_required_blocks(resources):
    """Make sub-fields inside mutually-exclusive provider blocks optional.

    Resources like castai_edge_location have cloud-specific blocks (aws, gcp,
    oci, custom) where only one is needed depending on the cloud provider.
    The Terraform schema marks sub-fields inside each block as required, but
    since the blocks themselves are optional and mutually exclusive, Upjet
    incorrectly generates CEL validation rules requiring all of them.

    This function converts required sub-fields inside those blocks to optional
    so that the CRD allows creating resources with only one provider block.
    """
    rules = {
        "castai_edge_location": ["aws", "gcp", "oci"],
    }

    relaxed = []
    for res_name, block_names in rules.items():
        if res_name not in resources:
            continue
        block_types = resources[res_name].get("block", {}).get("block_types", {})
        for bname in block_names:
            if bname not in block_types:
                continue
            inner_attrs = block_types[bname].get("block", {}).get("attributes", {})
            fields = []
            for fname, fattr in inner_attrs.items():
                if fattr.get("required"):
                    fattr.pop("required", None)
                    fattr["optional"] = True
                    fields.append(fname)
            if fields:
                relaxed.append(f"{res_name}.{bname}: {fields}")

    return relaxed


def main():
    schema_path = sys.argv[1] if len(sys.argv) > 1 else "config/schema.json"

    with open(schema_path) as f:
        schema = json.load(f)

    resources = schema["provider_schemas"]["registry.terraform.io/castai/castai"][
        "resource_schemas"
    ]

    changed = []
    for res_name, res in resources.items():
        block = res.get("block", {})
        nested = [k for k, v in block.get("attributes", {}).items() if "nested_type" in v]
        if nested:
            transform_block(block)
            changed.append(f"{res_name}: {nested}")

    relaxed = relax_conditionally_required_blocks(resources)

    with open(schema_path, "w") as f:
        json.dump(schema, f, indent=2)

    if changed:
        print(f"Transformed nested_type -> block_types in {len(changed)} resource(s):")
        for c in changed:
            print(f"  {c}")
    else:
        print("No nested_type attributes found, schema unchanged.")

    if relaxed:
        print(f"Relaxed conditionally-required sub-fields in {len(relaxed)} block(s):")
        for r in relaxed:
            print(f"  {r}")


if __name__ == "__main__":
    main()
