// FirstPageStyles.tsx

import styled from 'styled-components';

export const PageContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background-color: #fff;
  padding: 20px;
  transition: background-color 0.3s ease-in-out, transform 0.2s ease-in-out;
  box-sizing: border-box;
`;

export const FormContainer = styled.div`
  background: white;
  padding: 2rem;
  border-radius: 10px;
  width: 100%;
  max-width: 400px;
  box-sizing: border-box;
  margin: auto;
`;

export const Title = styled.h1`
  font-size: 3rem;
  margin-bottom: 1.5rem;
  color: #333;
  text-align: center;
`;

export const SubTitle = styled.p`
  font-size: 1.25rem;
  margin-bottom: 1.5rem;
  font-weight: bold;
  color: #333;
  text-align: center;
`;

export const Button1 = styled.button`
  width: 100%;
  height: 3rem;
  padding: 0.75rem;
  margin-bottom: 1rem;
  border: none;
  border-radius: 5px;
  background-color: #eede77;
  color: #333;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s ease-in-out, transform 0.2s ease-in-out;

  &:hover {
    background-color: #EBD961;
  }
  `;

export const Button2 = styled.button`
  width: 100%;
  padding: 0.75rem;
  margin-bottom: 1rem;
  border: 2px solid #eede77;
  background-color: #fff;
  border-radius: 5px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s ease-in-out, transform 0.2s ease-in-out;

  &:hover {
    background-color: #f9f9f9;
  }
`;