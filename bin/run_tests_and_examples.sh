# Run all tests
cd ../structures
echo "================================================="
echo "==== RUNNING ALL STRUCTURES TESTS/EXAMPLES...===="
echo "================================================="
go test ./...
cd -


# FIS size tests
cd ../evaluation
echo "================================================="
echo "==== RUNNING FREQUENT ITEMSETS SIZE TESTS...====="
echo "================================================="
go run fis_size-tests.go
cd -

# Skip lists complexity tests
cd ../evaluation
echo "================================================="
echo "==== RUNNING SKIP LIST PERFORMANCE TESTS...======"
echo "================================================="
go run skip-list_performance.go
cd -
