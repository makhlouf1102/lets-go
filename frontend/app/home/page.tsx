'use client'
import useSWR from "swr"
import { Problem, ProblemAccordion } from "@/components/ProblemAccordion"
import { LoadingState } from "@/components/LoadingState"
import { ErrorState } from "@/components/ErrorState"
import { InfoCard } from "@/components/InfoCard"

async function getProblems() {
    return fetch('http://localhost:8080/problems')
        .then(res => res.json())
        .catch(error => new Error("Failed to fetch problems" + String(error)))
}

export default function Home() {
    const { data, error, isLoading } = useSWR('/api/problems', getProblems)

    return (
        <div className="min-h-screen bg-base-200 py-12 px-4 sm:px-6 lg:px-8">
            <div className="max-w-7xl mx-auto">
                <div className="text-center mb-12">
                    <h1 className="text-4xl font-extrabold tracking-tight text-base-content sm:text-5xl mb-2">
                        Coding Challenges
                    </h1>
                    <p className="text-lg text-base-content/70">
                        Select a problem to start solving.
                    </p>
                </div>

                <div className="grid grid-cols-1 lg:grid-cols-4 gap-8 items-start">
                    {/* Left Column */}
                    <div className="space-y-8 lg:sticky lg:top-8">
                        <InfoCard title="How it's built">
                            <p>Built with a modern stack:</p>
                            <ul className="list-disc list-inside mt-2 space-y-1">
                                <li>Next.js 16 (Frontend)</li>
                                <li>Go (Backend)</li>
                                <li>DaisyUI & Tailwind</li>
                                <li>Dockerized</li>
                            </ul>
                        </InfoCard>

                        <InfoCard title="Who I am">
                            <p>Just a passionate developer who loves building cool things and solving interesting problems.</p>
                        </InfoCard>
                    </div>

                    {/* Center Column - Problems List */}
                    <div className="lg:col-span-2 space-y-6">
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

                    {/* Right Column */}
                    <div className="lg:sticky lg:top-8">
                        <InfoCard title="Why this project?">
                            <p>To practice full-stack development, experiment with new UI concepts, and create a platform for sharing coding challenges.</p>
                        </InfoCard>
                    </div>
                </div>
            </div>
        </div>
    )
}