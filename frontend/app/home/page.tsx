'use client'
import useSWR from "swr"
import { Problem, ProblemAccordion } from "@/components/ProblemAccordion"

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

                {isLoading && (
                    <div className="flex flex-col items-center justify-center py-20 space-y-4">
                        <span className="loading loading-infinity loading-lg text-primary scale-150"></span>
                        <p className="text-base-content/60 animate-pulse">Loading challenges...</p>
                    </div>
                )}

                {error && (
                    <div role="alert" className="alert alert-error shadow-lg animate-in fade-in slide-in-from-bottom-4 duration-500">
                        <svg xmlns="http://www.w3.org/2000/svg" className="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
                        <div>
                            <h3 className="font-bold text-lg">Oops! Something went wrong.</h3>
                            <div className="text-sm opacity-90">{error.message || "Failed to load problems. Please try again later."}</div>
                        </div>
                    </div>
                )}

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