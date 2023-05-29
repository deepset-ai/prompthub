from pathlib import Path

import yaml


def is_valid(prompt_file: Path) -> bool:
    if not prompt_file.exists():
        # We consider a non existing file to be valid
        return True

    try:
        data = yaml.safe_load(prompt_file.read_text())
    except:
        # Parsing failure means either the file is not yaml or it's malformed
        print(f"Failed parsing file: {prompt_file}")
        return False

    required_fields = {
        "name": str,
        "prompt_text": str,
        "description": str,
        "tags": list,
        "meta": dict,
        "version": str,
    }

    reports = []
    for key, expected_type in required_fields.items():
        if key not in data:
            reports.append(f"Required field {key} is missing")
            continue

        value = data[key]
        # Check the type of each field is the one we expect
        if type(value) is not expected_type:
            reports.append(f"Field {key} must be of type {expected_type}")
            continue

        # For list fields we have an extra check as we want to verify that they contain strings
        if expected_type is list and len(value) > 0 and type(value[0]) is not str:
            reports.append(f"Field {key} must be a list of strings")

    # Print a report if any field doesn't respect requirements
    if reports:
        print(f"File not valid: {prompt_file}")
        for report in reports:
            print(f"  - {report}")

    # We got reports, the file must not be valid
    return len(reports) == 0


if __name__ == "__main__":
    import sys
    import argparse

    parser = argparse.ArgumentParser(
        description="Utility script to validate prompts",
    )
    parser.add_argument("files", nargs="+")
    args = parser.parse_args()

    files = [Path(f) for f in args.files]

    print()
    any_not_valid = any([is_valid(f) for f in files])

    if any_not_valid:
        sys.exit(1)
