# How to Solve

* For Solving this problem you should map the characters to their pairs and use `stack` to push and pop.

## Solving Steps

1. Create a map from open brackets to close brackets.

2. Create a map from close brackets to open brackets.

3. Create a stack with a slice of items.

4. Create a push method to add brackets to the stack

5. Create a pop method to return the top item of the stack.

6. Create an instance of the stack.

7. Loop over the input string.

8. If bracket exists in the open map add it to stack.

9. If bracket exists in the close map:

    1. If stack is empty return false.

    2. Get the top of the stack.

    3. If top is not equal to the close map bracket return false.

10. If stack is empty return true.