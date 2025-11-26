'use client'
import useSWR from "swr"
import { Problem, ProblemAccordion } from "@/components/ProblemAccordion"
import { LoadingState } from "@/components/LoadingState"
import { ErrorState } from "@/components/ErrorState"

async function getProblems() {
    return fetch('http://localhost:8080/problems')
        .then(res => res.json())
        .catch(error => new Error("Failed to fetch problems" + String(error)))
}

export default function Home() {
    const { data, error, isLoading } = useSWR('/api/problems', getProblems)

    return (
        <div className="min-h-screen bg-base-200 flex flex-col items-center py-12 px-4 sm:px-6 lg:px-8">
            <div className="w-full max-w-3xl space-y-8">
                <div className="text-center">
                    <h1 className="text-4xl font-extrabold tracking-tight text-base-content sm:text-5xl mb-2">
                        Coding Challenges
                    </h1>
                    <p className="text-lg text-base-content/70">
                        Select a problem to start solving.
                    </p>
                </div>

                {isLoading && <LoadingState />}

                {error && <ErrorState message={error.message} />}

                {data && (
                    <div className="flex flex-col gap-4">
                        {data.data.map((problem: Problem) => (
                            <ProblemAccordion key={problem.id} problem={problem} />
                        ))}
                    </div>
                )}
            </div>
        </div>
    )
}