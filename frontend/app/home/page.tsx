'use client'
import useSWR from "swr"

type Problem = {
    id: number;
    title: string;
    description: string;
    difficulty: string;
}

async function getProblems() {
    return fetch('http://localhost:8080/problems')
        .then(res => res.json())
        .catch(error => new Error("Failed to fetch problems" + String(error)))
}

const getDifficultyColor = (difficulty: string) => {
    switch (difficulty.toLowerCase()) {
        case 'easy': return 'badge-success';
        case 'medium': return 'badge-warning';
        case 'hard': return 'badge-error';
        default: return 'badge-ghost';
    }
};

export default function Home() {
    const { data, error, isLoading } = useSWR('/api/problems', getProblems)

    return (
        <div>
            {isLoading && <p>Loading...</p>}
            {error && <p>Error: {error.message}</p>}
            {data && <pre>{JSON.stringify(data, null, 2)}</pre>}
        </div>
    )
}