import React from "react";

interface CardProps {
  children: React.ReactNode;
}

const Card: React.FC<CardProps> = ({ children }) => {
  return (
    <>
      <div className="bg-white rounded-xl shadow-lg p-8 w-11/12 max-w-md text-center flex flex-col justify-between min-h-[400px] mt-4">
        {children}
      </div>
      <div className="text-center mt-4 text-gray-500">
        <p>
          Check out the code on{" "}
          <a
            href="https://github.com/martishin/react-golang-user-login-oauth"
            target="_blank"
            className="underline text-gray-500 hover:text-gray-700"
          >
            GitHub
          </a>
        </p>
      </div>
    </>
  );
};

export default Card;
