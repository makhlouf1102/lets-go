import React from 'react';

interface ErrorStateProps {
    message?: string;
}

export const ErrorState: React.FC<ErrorStateProps> = ({ message }) => {
    return (
        <div role="alert" className="alert alert-error shadow-lg animate-in fade-in slide-in-from-bottom-4 duration-500">
            <svg xmlns="http://www.w3.org/2000/svg" className="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
            <div>
                <h3 className="font-bold text-lg">Oops! Something went wrong.</h3>
                <div className="text-sm opacity-90">{message || "Failed to load problems. Please try again later."}</div>
            </div>
        </div>
    );
};
