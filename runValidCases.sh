# Run normal tests
echo "🚀 Run normal test cases!"

echo "📑 Example 00:"
go run . "examples/example00.txt" 
echo "➡️ Took steps: $(go run . "examples/example00.txt" | grep -c '^L'), expected <=6"
echo ""

echo "📑 Example 01:"
go run . "examples/example01.txt" 
echo "➡️ Took steps: $(go run . "examples/example01.txt" | grep -c '^L'), expected <=8"
echo ""

echo "📑 Example 02:"
go run . "examples/example02.txt" 
echo "➡️ Took steps: $(go run . "examples/example02.txt" | grep -c '^L'), expected <=11"
echo ""

echo "📑 Example 03:"
go run . "examples/example03.txt" 
echo "➡️ Took steps: $(go run . "examples/example03.txt" | grep -c '^L'), expected <=6"
echo ""

echo "📑 Example 04:"
go run . "examples/example04.txt" 
echo "➡️ Took steps: $(go run . "examples/example04.txt" | grep -c '^L'), expected <=6"
echo ""

echo "📑 Example 05:"
go run . "examples/example05.txt" 
echo "➡️ Took steps: $(go run . "examples/example05.txt" | grep -c '^L'), expected <=8"
echo ""