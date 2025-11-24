
'use client'
import { ResultAsync } from "neverthrow"
import useSWR from "swr"

type Problem = {
    id: number;
    title: string;
    description: string;
    category: string;
    difficulty: string;
}

function getProblems() {
    return ResultAsync.fromPromise(fetch('http://localhost:8080/problems')
        .then(res => res.json()),
        (error) => new Error("Failed to fetch problems" + String(error)))
}

export default function Home() {
    const { data, error, isLoading } = useSWR('/api/problems', getProblems)

    if (isLoading) return <div>Loading...</div>
    if (error) return <div>Error: {error.message}</div>

    return (
        <div>

        </div>
    )
}