
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

async function getProblems() {
    return fetch('http://localhost:8080/problems')
        .then(res => res.json())
        .catch(error => new Error("Failed to fetch problems" + String(error)))
}

export default function Home() {
    const { data, error, isLoading } = useSWR('/api/problems', getProblems)

    return (
        <div>
            {isLoading && <p>Loading...</p>}
            {error && <p>Error: {error.message}</p>}
            {data && <ul>
                {data.data.map((problem: Problem) => (
                    <li key={problem.id}>{problem.title}</li>
                ))}
            </ul>}
        </div>
    )
}