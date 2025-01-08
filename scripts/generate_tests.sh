# First remove all existing test files
find test/v2025_01 -type f -name "*_test.go" -delete

# Then create new test files
index=0
find pkg/v2025_01/payload -type f -name "*.go" | while read -r file; do
    # Skip if it's already a test file
    if [[ "$file" == *_test.go ]]; then
        continue
    fi

    # Extract type name from New$Type function
    type=$(grep -o "func New[A-Za-z0-9_]*" "$file" | sed 's/func New//')

    if [ ! -z "$type" ]; then
        # Create test file name in the test directory
        filename=$(basename "$file")
        test_file="test/v2025_01/${filename%.go}_test.go"

        # Ensure the test directory exists
        mkdir -p test/v2025_01

        echo "package v2025_01__test

import (
    \"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/payload\"
    \"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook\"
    \"github.com/sadsciencee/shopify-webhooks-go/test\"
    \"testing\"
)

func cb${type}(t *testing.T, w *webhook.UnknownWebhook) error {
    webhook, err := payload.New$type(w)
    if err != nil {
        t.Fatalf(\"failed to parse $type: %v\", err)
    }
    _, err = webhook.GetData()
    if err != nil {
        t.Fatalf(\"failed to get data: %v\", err)
    }

    return nil
}

func Test${type}(t *testing.T) {
    t.Parallel()
    test.WebhookHelper(t, webhook.$type, cb${type}, $index)
}" > "$test_file"

        echo "Created test file for $type: $test_file"
        index=$((index + 1))
    fi
done