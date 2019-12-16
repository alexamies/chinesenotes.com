export class HrefVariableParser {
    getHrefVariable(href, name) {
        if (!href.includes("?")) {
            console.log("getHrefVariable: href does not include ? ", href);
            return "";
        }
        const path = href.split("?");
        const parts = path[1].split("&");
        for (let i = 0; i < parts.length; i += 1) {
            const p = parts[i].split("=");
            if (decodeURIComponent(p[0]) == name) {
                return decodeURIComponent(p[1]);
            }
        }
        console.log(`getHrefVariable: ${name} not found`);
        return "";
    }
}
//# sourceMappingURL=HrefVariableParser.js.map