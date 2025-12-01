"use client"
import Editor from '@monaco-editor/react';

export default function ProblemPage() {
    return <Editor height="90vh" defaultLanguage="javascript" defaultValue="// some comment" />;
}