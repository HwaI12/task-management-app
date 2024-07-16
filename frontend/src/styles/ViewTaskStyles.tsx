import styled from 'styled-components';

export const MarkdownPreview = styled.div`
    width: 80%;
    padding: 20px;
    border-radius: 8px;
    line-height: 1.6;
    color: #333;

    h2 {
        margin-top: 0;
    }

    h3 {
        margin-bottom: 0.5em;
    }

    p {
        margin: 0.5em 0;
    }

    ul {
        padding-left: 20px;
    }

    li {
        margin-bottom: 0.5em;
    }

    pre {
        background-color: #333;
        color: #f7f7f7;
        padding: 10px;
        border-radius: 5px;
        overflow-x: auto;
    }

    code {
        background-color: #eaeaea;
        padding: 2px 4px;
        border-radius: 4px;
    }
`;

export const TaskContainer = styled.div`
    display: flex;
    justify-content: center;
    padding: 20px;
    box-sizing: border-box;

    h2 {
        font-size: 1.5em;
        margin-bottom: 20px;
    }
`;
