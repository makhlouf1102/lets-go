import React from 'react';

interface InfoCardProps {
    title: string;
    children: React.ReactNode;
}

export const InfoCard: React.FC<InfoCardProps> = ({ title, children }) => {
    return (
        <div className="card bg-base-100 shadow-xl border-2 border-base-content/10 hover:border-base-content/30 transition-all duration-300 h-full">
            <div className="card-body p-6">
                <h2 className="card-title text-2xl mb-2">{title}</h2>
                <div className="text-lg leading-relaxed opacity-90">
                    {children}
                </div>
            </div>
        </div>
    );
};
