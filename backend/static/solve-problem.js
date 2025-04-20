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
    window.editor = monaco.editor.create(document.getElementById('editor'), {
        value: sourceCode,
        language: currentProggramingLanguage,
        automaticLayout: true,
        padding: { top: 5, right: 5, bottom: 5, left: 5 },
        overviewRulerLanes: 0,
        overviewRulerBorder: false,
        theme: 'vs-dark',
    });
});

var programmingLanguagesSelect = document.getElementById("programming-language")

programmingLanguagesSelect.addEventListener("sl-change", (e) => {
    currentProggramingLanguage = e.target.value
    alert(e.target.value)
    proxyApiService.getProblemCode(currentProggramingLanguage, problemId)
        .then((data) => {
            sourceCode = data
            window.editor.value = sourceCode
        })
        .catch((error) => {
            console.log('Error', error)
        })
})

document.getElementById("submit-button").addEventListener("click", async (e) => {
    var code = window.editor.getValue()
    var language = currentProggramingLanguage

    proxyApiService.runCode(language, code).then((data) => {
        console.log(data)
    })
        .catch((error) => {
            console.error("Error", error)
        })
})

