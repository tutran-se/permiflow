#!/bin/bash

LOG_FILE="permiflow_test_output.log"
BIN="go run ."

echo "🔍 Starting Permiflow v0.2.0 Test Run" > $LOG_FILE
echo "Timestamp: $(date)" >> $LOG_FILE
echo "----------------------------------" >> $LOG_FILE

function run_test() {
  echo -e "\n$1" | tee -a $LOG_FILE
  echo "Command: $2" | tee -a $LOG_FILE
  eval $2 >> $LOG_FILE 2>&1
  echo "----------------------------------" >> $LOG_FILE
}

# 1. CLI flags
run_test "TEST: Help Output" "$BIN --help"
run_test "TEST: Version Output" "$BIN --version"
run_test "TEST: Emoji Disabled with --plain" "$BIN --dry-run --plain"
run_test "TEST: Emoji Disabled with env var" "PERMIFLOW_NO_EMOJI=true $BIN --dry-run"

# 2. Namespace behaviors
run_test "TEST: Valid Namespace (default)" "$BIN --namespace=default --dry-run"
run_test "TEST: Non-existent Namespace (nope)" "$BIN --namespace=nope --dry-run"
run_test "TEST: Empty Namespace Scan (cluster-wide)" "$BIN --dry-run"

# 3. Output generation
run_test "TEST: Markdown + CSV Output" "$BIN --out-dir=tmp/out --prefix=testscan"
run_test "TEST: Dry Run (No File Write)" "$BIN --dry-run"
run_test "TEST: Output to nested directory" "$BIN --out-dir=tmp/reports/test01"

# 4. Risk logic samples (manual setup expected)
run_test "TEST: Wildcard Verb Detection" "$BIN --namespace=default --dry-run"
run_test "TEST: Secrets Access Detection" "$BIN --namespace=default --dry-run"
run_test "TEST: cluster-admin Detection" "$BIN --dry-run"

# 5. Broken input
run_test "TEST: Bad kubeconfig path" "$BIN --kubeconfig=/bad/path --dry-run"

# Summary
echo -e "\n✅ Test run complete. Output saved to $LOG_FILE."
