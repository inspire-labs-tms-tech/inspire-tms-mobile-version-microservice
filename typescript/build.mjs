// @ts-check
import {generateApi} from "swagger-typescript-api";
import * as path from "node:path";
import {fileURLToPath} from 'node:url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);


/**
 * @typedef {import("swagger-typescript-api").GenerateApiParams} GenerateApiParams
 */

/** @type {GenerateApiParams["hooks"]["onFormatRouteName"]} */
const onFormatRouteName = ({route}) => {
    const parts = route
        .split("/")
        .filter(v => !!v) // remove empty splits

    if(parts.length === 1) return "Get"
    else return parts.map(part => part.charAt(0).toUpperCase() + part.slice(1)).join("") // Capitalize the first letter of each part
}

/** @type {GenerateApiParams} */
await generateApi({
    name: "index.ts",
    output: path.resolve(__dirname),
    input: path.resolve(__dirname, "../OpenAPI.yaml"),
    apiClassName: "MobileVersionMicroservice",
    httpClientType: "fetch",
    hooks: {
        onFormatRouteName
    },
});