import styled from 'styled-components';

export const ContentContainer = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 20px;
    margin-left: 60px;
    text-align: left;
`;

export const Form = styled.form`
    width: 100%;
    max-width: 600px;
    margin-top: 20px;
`;

export const FormGroup = styled.div`
    margin-bottom: 30px;
`;

export const TitleLabel = styled.label`
    display: block;
    font-weight: bold;
    margin-bottom: 5px;
    font-align: left;
    font-size: 2rem; /* タイトル用の大きさ */
`;

export const SubtitleLabel = styled.label`
    display: block;
    font-weight: bold;
    margin-bottom: 5px;
    font-size: 1.5rem; /* タイトル以外用の大きさ */
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

export const Select = styled.select`
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
    border-radius: 5px;: 5px;
    cursor: pointer;
    transition: background-color 0.3s ease-in-out, transform 0.2s ease-in-out;
    &:hover {
        background-color: #EBD961;
    }
`;

export const MarkdownPreview = styled.div`
    background-color: #fff;
    padding: 10px;
    margin-top: 20px;
    width: 100%;
    max-width: 600px;
`;

export const InputIconWrapper = styled.div`
    display: flex;
    align-items: center;
    background-color: #f0f0f0;
    border-radius: 5px;
    padding: 10px;
    margin-bottom: 20px;
`;

export const InputIcon = styled.span`
    margin-right: 10px;
    font-size: 1.2rem;
`;

export const StyledTextarea = styled(Textarea)`
    height: 100px;
`;

// interface ToggleButtonProps {
//     active: boolean;
// }

export const ButtonGroup = styled.div`
    display: flex;
    justify-content: center;
    margin-bottom: 20px;
`;

// export const ToggleButton = styled.button<ToggleButtonProps>`
//     border: none;
//     padding: 10px 20px;
//     margin: 0 10px;
//     cursor: pointer;
//     border-radius: 5px;
//     font-size: 1rem;
//     &:hover {
//         background-color: ${props => props.active ? '#0056b3' : '#bbb'};
//     }
// `;
