"use client"
import Editor from '@monaco-editor/react';
import { Panel, PanelGroup, PanelResizeHandle } from 'react-resizable-panels';
import { useState, useEffect, useRef, RefObject } from 'react';
import { editor } from 'monaco-editor';
import { useParams, useRouter } from 'next/navigation';

interface JudgeResponse {
    stdout: string;
    stderr: string;
    time: string;
    memory: number;
    token: string;
    status: {
        id: number;
        description: string;
    }
}

interface RunCodeResponse {
    data: JudgeResponse;
    message: string;
}

interface Problem {
    id: number;
    title: string;
    description: string;
    signature: string;
    difficulty: string;
}

interface ProblemApiResponse {
    message: string;
    data: Problem;
}

async function runCode(editorRef: RefObject<editor.IStandaloneCodeEditor | null>): Promise<RunCodeResponse> {
    try {
        const response = await fetch('http://localhost:8080/code/run', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                source_code: editorRef?.current?.getValue() || '',
                language_id: 63,
            }),
        });
        const data = await response.json();
        return data as RunCodeResponse;
    } catch (error) {
        alert(String(error));
        console.error('Error running code:', error);
        return { data: { stdout: '', stderr: '', time: '', memory: 0, token: '', status: { id: 0, description: '' } }, message: String(error) };
    }
}

async function fetchProblem(id: string): Promise<Problem | null> {
    try {
        const response = await fetch(`http://localhost:8080/problems/${id}`);
        if (!response.ok) {
            throw new Error(`Failed to fetch problem: ${response.statusText}`);
        }
        const data: ProblemApiResponse = await response.json();
        return data.data;
    } catch (error) {
        console.error('Error fetching problem:', error);
        return null;
    }
}

export default function ProblemPage() {
    const params = useParams();
    const router = useRouter();
    const id = params.id as string;

    const [direction, setDirection] = useState<'horizontal' | 'vertical'>('horizontal');
    const [problem, setProblem] = useState<Problem | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const handleResize = () => {
            if (window.innerWidth < 1024) {
                setDirection('vertical');
            } else {
                setDirection('horizontal');
            }
        };

        handleResize(); // Initial check
        window.addEventListener('resize', handleResize);
        return () => window.removeEventListener('resize', handleResize);
    }, []);

    useEffect(() => {
        async function loadProblem() {
            setLoading(true);
            setError(null);
            const problemData = await fetchProblem(id);
            if (problemData) {
                setProblem(problemData);
            } else {
                setError('Failed to load problem');
            }
            setLoading(false);
        }

        if (id) {
            loadProblem();
        }
    }, [id]);

    const editorRef = useRef<editor.IStandaloneCodeEditor | null>(null);
    const [output, setOutput] = useState('');

    function handleEditorDidMount(editor: editor.IStandaloneCodeEditor, _: any) {
        editorRef.current = editor;
    }

    async function showValue() {
        const result = await runCode(editorRef);
        console.log(result);
        setOutput(result.data.stdout);
    }

    if (loading) {
        return (
            <div className="min-h-screen bg-base-200 p-8 flex items-center justify-center">
                <div className="text-center">
                    <span className="loading loading-spinner loading-lg"></span>
                    <p className="mt-4 text-lg">Loading problem...</p>
                </div>
            </div>
        );
    }

    if (error || !problem) {
        return (
            <div className="min-h-screen bg-base-200 p-8 flex items-center justify-center">
                <div className="card w-full max-w-md bg-base-100 shadow-xl">
                    <div className="card-body text-center">
                        <h2 className="card-title text-error justify-center">Error</h2>
                        <p>{error || 'Problem not found'}</p>
                        <div className="card-actions justify-center mt-4">
                            <button className="btn btn-primary" onClick={() => router.back()}>
                                Go Back
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        );
    }

    return (
        <div className="min-h-screen bg-base-200 p-8 flex items-center justify-center">
            <div className="card w-full max-w-[95vw] h-[90vh] bg-base-100 shadow-xl overflow-hidden border border-base-300 rounded-2xl flex flex-col">
                {/* Header */}
                <div className="bg-base-100 border-b border-base-300 p-3 flex items-center gap-4 shrink-0">
                    <button className="btn btn-ghost btn-circle btn-sm" onClick={() => router.back()}>
                        <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10 19l-7-7m0 0l7-7m-7 7h18" />
                        </svg>
                    </button>
                    <h1 className="font-bold text-lg">{problem.title}</h1>
                    <span className={`badge ${problem.difficulty === 'easy' ? 'badge-success' :
                            problem.difficulty === 'medium' ? 'badge-warning' :
                                'badge-error'
                        }`}>
                        {problem.difficulty}
                    </span>
                </div>

                <div className="flex-1 overflow-hidden">
                    <PanelGroup direction={direction}>
                        {/* Left Panel */}
                        <Panel defaultSize={33} minSize={20} className="flex flex-col bg-base-100">
                            <div className="p-6 overflow-y-auto flex-1">
                                <div className="prose max-w-none">
                                    <h2 className="text-2xl font-bold mb-4">Description</h2>
                                    <div className="text-base-content/80 mb-6 whitespace-pre-wrap">
                                        {problem.description}
                                    </div>
                                </div>
                            </div>
                        </Panel>

                        <PanelResizeHandle className={`bg-base-300 transition-colors hover:bg-primary/50 ${direction === 'horizontal' ? 'w-1 cursor-col-resize' : 'h-1 cursor-row-resize'}`} />

                        {/* Right Panel */}
                        <Panel minSize={30} className="flex flex-col">
                            <PanelGroup direction="vertical">
                                {/* Code Section */}
                                <Panel defaultSize={70} minSize={20} className="flex flex-col">
                                    <div className="bg-base-200 px-4 py-2 text-xs font-bold uppercase tracking-wider text-base-content/60 border-b border-base-300 select-none shrink-0">
                                        Code
                                    </div>
                                    <div className="flex-1 relative">
                                        <Editor
                                            height="100%"
                                            defaultLanguage="javascript"
                                            defaultValue={problem.signature || "// Write your solution here"}
                                            options={{
                                                minimap: { enabled: false },
                                                fontSize: 14,
                                                scrollBeyondLastLine: false,
                                                padding: { top: 16, bottom: 16 }
                                            }}
                                            onMount={handleEditorDidMount}
                                        />
                                    </div>
                                </Panel>

                                <PanelResizeHandle className="h-1 bg-base-300 cursor-row-resize transition-colors hover:bg-primary/50" />

                                {/* Response Section */}
                                <Panel defaultSize={30} minSize={10} className="flex flex-col bg-base-100">
                                    <div className="bg-base-200 px-4 py-2 flex justify-between items-center border-b border-base-300 shrink-0">
                                        <span className="text-xs font-bold uppercase tracking-wider text-base-content/60 select-none">Response</span>
                                        <button className="btn btn-primary btn-sm px-6" onClick={showValue}>Run</button>
                                    </div>
                                    <div className="p-4 font-mono text-sm overflow-y-auto flex-1 bg-base-100 text-base-content">
                                        <span className="opacity-50 italic">{output}</span>
                                    </div>
                                </Panel>
                            </PanelGroup>
                        </Panel>
                    </PanelGroup>
                </div>
            </div>
        </div>
    );
}
