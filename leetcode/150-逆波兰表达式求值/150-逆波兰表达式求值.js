/**
 * @param {string[]} tokens
 * @return {number}
 */
var evalRPN = function (tokens) {
    const stack = [];
    const n = tokens.length;
    for (let i = 0; i < n; i++) {
        if (tokens[i] === "+" || tokens[i] === "*" || tokens[i] === "-" || tokens[i] === "/") {
            const num1 = stack.pop();
            const num2 = stack.pop();
            const operator = tokens[i];
            if (operator === "*") {
                stack.push(num1 * num2);
            } else if (operator === "+") {
                stack.push(num1 + num2);
            } else if (operator === "-") {
                stack.push(num2 - num1);
            } else if (operator === "/") {
                stack.push(num2 / num1 > 0 ? Math.floor(num2 / num1) : Math.ceil(num2 / num1));
            }
        } else {
            stack.push(parseInt(tokens[i]));
        }
    }
    return stack.pop();
};