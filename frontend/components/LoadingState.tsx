import React from 'react';

export const LoadingState: React.FC = () => {
    return (
        <div className="flex flex-col items-center justify-center py-20 space-y-4">
            <span className="loading loading-infinity loading-lg text-primary scale-150"></span>
            <p className="text-base-content/60 animate-pulse">Loading challenges...</p>
        </div>
    );
};
