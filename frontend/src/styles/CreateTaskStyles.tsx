import styled from 'styled-components';
import { StylesConfig } from 'react-select';

export const CreateTaskContainer = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: left;
    padding: 20px;

    @media (max-width: 768px) {
        padding: 30px;
        margin-left: 0;
    }
`;

export const Form = styled.form`
    width: 100%;
    max-width: 600px;
    margin-top: 20px;

    @media (max-width: 768px) {
        max-width: 100%;
    }
`;

export const FormGroup = styled.div`
    margin-bottom: 30px;
    width: 100%;
`;

export const TitleLabel = styled.label`
    display: block;
    font-weight: bold;
    margin-bottom: 5px;
    font-size: 2rem;

    @media (max-width: 768px) {
        font-size: 1.5rem;
    }
`;

export const SubtitleLabel = styled.label`
    display: block;
    font-weight: bold;
    margin-bottom: 5px;
    font-size: 1.5rem;

    @media (max-width: 768px) {
        font-size: 1.2rem;
    }
`;

export const Input = styled.input`
    width: 100%;
    padding: 10px;
    font-size: 1rem;
    border: 1px solid #f0f0f0;
    border-radius: 5px;
    background-color: #f0f0f0;
    outline: none;
`;

export const Textarea = styled.textarea`
    width: 100%;
    padding: 10px;
    font-size: 1rem;
    border: 1px solid #f0f0f0;
    border-radius: 5px;
    background-color: #f0f0f0;
    outline: none;
`;

export const ErrorMessage = styled.p`
    color: red;
    font-weight: bold;
`;

export const Button = styled.button`
    width: 100%;
    padding: 15px;
    font-size: 1.2rem;
    color: #333;
    background-color: #eede77;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    transition: background-color 0.3s ease-in-out, transform 0.2s ease-in-out;

    &:hover {
        background-color: #EBD961;
    }

    @media (max-width: 768px) {
        font-size: 1rem;
    }
`;

export const MarkdownPreview = styled.div`
    background-color: #fff;
    padding: 10px;
    margin-top: 20px;
    width: 100%;
    max-width: 600px;

    @media (max-width: 768px) {
        max-width: 100%;
    }
`;

export const InputIconWrapper = styled.div`
    display: flex;
    align-items: center;
    background-color: #f0f0f0;
    border-radius: 5px;
    padding: 10px;
    margin-bottom: 20px;
    width: 100%;
    box-sizing: border-box;
`;

export const InputIcon = styled.span`
    margin-right: 10px;
    font-size: 1.2rem;
`;

export const SelectGroup = styled.div`
    width: 100%;
`;

export const StyledTextarea = styled(Textarea)`
    height: 100px;
`;

export const ButtonGroup = styled.div`
    display: flex;
    justify-content: center;
    margin-bottom: 20px;
`;

export const customStyles: StylesConfig = {
    control: (provided) => ({
        ...provided,
        fontSize: '1rem',
        border: '1px solid #f0f0f0',
        borderRadius: '5px',
        backgroundColor: '#f0f0f0',
        outline: 'none',
        appearance: 'none',
        position: 'relative',
        zIndex: 1,
        boxSizing: 'border-box',
    }),
    valueContainer: (provided) => ({
        ...provided,
        boxSizing: 'border-box',
    }),
    menu: (provided) => ({
        ...provided,
        zIndex: 9999,
    }),
    menuPortal: (base) => ({ ...base, zIndex: 9999 }),
    option: (provided) => ({
        ...provided,
        whiteSpace: 'nowrap',
        overflow: 'hidden',
        textOverflow: 'ellipsis',
    }),
};
