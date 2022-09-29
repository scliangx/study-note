/**
 * @param {string} path
 * @return {string}
 */

// 71-简化路径
var simplifyPath = function (path) {
    const paths = path.split("/");
    const stack = [];
    for (let i = 0; i < paths.length; i++) {
        if (paths[i] === "..") {
            if (stack.length > 0) {
                stack.pop();
            }
        } else if (paths[i] != "" && paths[i] != ".") {
            stack.push(paths[i]);
        }
    }
    return "/" + stack.join("/");
};