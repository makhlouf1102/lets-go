"use client"
import Editor from '@monaco-editor/react';

export default function ProblemPage() {
    return (
        <div className="min-h-screen bg-base-200 p-8 flex items-center justify-center">
            <div className="card w-full max-w-[95vw] h-[90vh] bg-base-100 shadow-xl overflow-hidden border border-base-300 rounded-2xl flex flex-col">
                {/* Header */}
                <div className="bg-base-100 border-b border-base-300 p-3 flex items-center gap-4">
                    <button className="btn btn-ghost btn-circle btn-sm">
                        <img src="https://placehold.co/24x24" alt="Back" className="w-6 h-6 rounded-full" />
                    </button>
                    <h1 className="font-bold text-lg">Problem Workspace</h1>
                </div>

                <div className="flex flex-1 overflow-hidden">
                    {/* Left Panel */}
                    <div className="w-1/3 min-w-[300px] border-r border-base-300 flex flex-col bg-base-100">
                        <div className="p-6 overflow-y-auto flex-1">
                            <div className="prose">
                                <h2 className="text-2xl font-bold mb-4">Description</h2>
                                <p className="text-base-content/80 mb-6">
                                    This is a placeholder for the problem description. In a real application, this would contain the detailed problem statement, examples, and constraints.
                                </p>

                                <div className="divider my-6"></div>

                                <h3 className="text-xl font-bold mb-3">Instruction</h3>
                                <p className="text-base-content/80">
                                    1. Read the problem description carefully.<br />
                                    2. Write your solution in the code editor on the right.<br />
                                    3. Click the "Run" button to test your code.<br />
                                    4. Check the output in the response section.
                                </p>
                            </div>
                        </div>
                    </div>

                    {/* Right Panel */}
                    <div className="flex-1 flex flex-col min-w-0">
                        {/* Code Section */}
                        <div className="flex-1 flex flex-col min-h-0">
                            <div className="bg-base-200 px-4 py-2 text-xs font-bold uppercase tracking-wider text-base-content/60 border-b border-base-300 select-none">
                                Code
                            </div>
                            <div className="flex-1 relative">
                                <Editor
                                    height="100%"
                                    defaultLanguage="javascript"
                                    defaultValue="// Write your solution here"
                                    options={{
                                        minimap: { enabled: false },
                                        fontSize: 14,
                                        scrollBeyondLastLine: false,
                                        padding: { top: 16, bottom: 16 }
                                    }}
                                />
                            </div>
                        </div>

                        {/* Response Section */}
                        <div className="h-1/3 flex flex-col border-t border-base-300 bg-base-100">
                            <div className="bg-base-200 px-4 py-2 flex justify-between items-center border-b border-base-300">
                                <span className="text-xs font-bold uppercase tracking-wider text-base-content/60 select-none">Response</span>
                                <button className="btn btn-primary btn-sm px-6">Run</button>
                            </div>
                            <div className="p-4 font-mono text-sm overflow-y-auto flex-1 bg-base-100 text-base-content">
                                <span className="opacity-50 italic">Output will appear here after running the code...</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}