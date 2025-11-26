import React from 'react';

export type Problem = {
    id: number;
    title: string;
    description: string;
    difficulty: string;
}

const getDifficultyColor = (difficulty: string) => {
    switch (difficulty.toLowerCase()) {
        case 'easy': return 'badge-success';
        case 'medium': return 'badge-warning';
        case 'hard': return 'badge-error';
        default: return 'badge-ghost';
    }
};

interface ProblemAccordionProps {
    problem: Problem;
}

export const ProblemAccordion: React.FC<ProblemAccordionProps> = ({ problem }) => {
    return (
        <div
            className="collapse collapse-arrow bg-base-100 shadow-sm hover:shadow-md transition-all duration-300 rounded-2xl border border-base-200"
        >
            <input type="radio" name="my-accordion-2" />
            <div className="collapse-title text-xl font-medium flex items-center gap-3">
                <span className="flex-1">{problem.title}</span>
                <span className={`badge ${getDifficultyColor(problem.difficulty)} badge-sm uppercase font-bold tracking-wider`}>
                    {problem.difficulty}
                </span>
            </div>
            <div className="collapse-content">
                <div className="divider my-0 opacity-50"></div>
                <div className="pt-4 flex flex-col gap-4">
                    <p className="text-base-content/80 leading-relaxed">
                        {problem.description}
                    </p>
                    <div className="flex justify-end pt-2">
                        <button className="btn btn-primary btn-sm px-6 shadow-sm hover:shadow-md transition-all">
                            Solve Challenge
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
};
