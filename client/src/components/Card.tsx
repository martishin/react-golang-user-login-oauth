import React from "react";

interface CardProps {
  children: React.ReactNode;
}

const Card: React.FC<CardProps> = ({ children }) => {
  return (
    <div className="bg-white rounded-xl shadow-lg p-8 w-11/12 max-w-md text-center flex flex-col justify-between min-h-[400px]">
      {children}
    </div>
  );
};

export default Card;
