import {defineConfig} from "tsup";

export default defineConfig({
    entry: ["./index.ts"],
    format: ["cjs", "esm"],
    outDir: ".",
    dts: true,
    splitting: true,
    sourcemap: true,
    clean: true,
    bundle: false,
    external: []
})