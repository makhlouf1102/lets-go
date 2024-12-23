import { proxyApiService } from "./ProxyApiService.js";

require.config({
    paths: {
        vs: 'https://cdnjs.cloudflare.com/ajax/libs/monaco-editor/0.39.0/min/vs',
    },
});

const url = new URL(window.location.href)
const problemId = url.searchParams.get("problemId")
var currentProggramingLanguage = "javascript"
var sourceCode = ""
proxyApiService.getProblemCode(currentProggramingLanguage, problemId)
    .then((data) => {
        sourceCode = data
    })
    .catch((error) => {
        console.log('Error', error)
    })


require(['vs/editor/editor.main'], () => {
    monaco.editor.create(document.getElementById('editor'), {
        value: sourceCode,
        language: currentProggramingLanguage,
        automaticLayout: true,
        padding: { top: 5, right: 5, bottom: 5, left: 5 },
        overviewRulerLanes: 0,
        overviewRulerBorder: false,
        theme: 'vs-dark',
    });
});