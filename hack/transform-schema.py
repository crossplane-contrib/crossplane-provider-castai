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
            converted = {
                "block": {
                    "attributes": inner_block["attributes"],
                    "description_kind": "plain",
                },
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

    with open(schema_path, "w") as f:
        json.dump(schema, f, indent=2)

    if changed:
        print(f"Transformed nested_type -> block_types in {len(changed)} resource(s):")
        for c in changed:
            print(f"  {c}")
    else:
        print("No nested_type attributes found, schema unchanged.")


if __name__ == "__main__":
    main()
