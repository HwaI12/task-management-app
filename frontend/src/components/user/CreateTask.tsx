import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate, useParams, Navigate } from 'react-router-dom';
import Sidebar from '../common/Sidebar';
import { ContentContainer, Form, FormGroup, TitleLabel, SubtitleLabel, Input, Select, ErrorMessage, Button, InputIconWrapper, StyledTextarea } from '../../styles/CreateTaskStyles';
// import ReactMarkdown from 'react-markdown';
// MarkdownPreview,ButtonGroup, ToggleButton

// ユーザーが認証されているか確認する関数
const isAuthenticated = (): boolean => {
    return localStorage.getItem('authToken') !== null;
};

const CreateTask: React.FC = () => {
    const [errorMessage, setErrorMessage] = useState('');
    const [title, setTitle] = useState('');
    const [deadline, setDeadline] = useState('');
    const [priority, setPriority] = useState('高');
    const [status, setStatus] = useState('未着手');
    const [purpose, setPurpose] = useState('');
    const [steps, setSteps] = useState('');
    const [memo, setMemo] = useState('');
    const [remarks, setRemarks] = useState('');
    // const [isEditing, setIsEditing] = useState(true);
    
    const { userId } = useParams<{ userId: string }>();

    const navigate = useNavigate();

    if (!isAuthenticated()) {
        return <Navigate to="/Signin" replace />;
    }

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setErrorMessage('');
        try {
            const response = await axios.post('http://localhost:8000/api/tasks', {
                user_id: userId,
                title,
                deadline,
                priority,
                status,
                purpose,
                steps,
                memo,
                remarks,
            });

            console.log('登録成功:', response.data);
            navigate(`/${userId}`);
        } catch (error) {
            if (axios.isAxiosError(error)) {
                console.error('登録エラー:', error.response?.data || error.message);
                setErrorMessage(error.response?.data?.message || '登録に失敗しました。');
            } else {
                console.error('予期しないエラー:', error);
                setErrorMessage('予期しないエラーが発生しました。');
            }
        }
    };

    return (
        <div>
            <Sidebar />
            <ContentContainer>
                <h2>タスク追加ページ</h2>
                <p>ユーザー {userId}!</p>

                {/* <ButtonGroup>
                    <ToggleButton active={isEditing} onClick={() => setIsEditing(true)}>Edit</ToggleButton>
                    <ToggleButton active={!isEditing} onClick={() => setIsEditing(false)}>Preview</ToggleButton>
                </ButtonGroup> */}

                {/* {isEditing ? ( */}
                    <Form onSubmit={handleSubmit}>
                        <FormGroup>
                            <TitleLabel>タイトル</TitleLabel>
                            <InputIconWrapper>
                                <Input
                                    type="text"
                                    value={title}
                                    onChange={(e) => setTitle(e.target.value)}
                                />
                            </InputIconWrapper>
                        </FormGroup>
                        <FormGroup>
                            <SubtitleLabel>目的</SubtitleLabel>
                            <InputIconWrapper>
                                <Input
                                    value={purpose}
                                    onChange={(e) => setPurpose(e.target.value)}
                                ></Input>
                            </InputIconWrapper>
                        </FormGroup>
                        <FormGroup>
                            <SubtitleLabel>優先度</SubtitleLabel>
                            <InputIconWrapper>
                                <Select value={priority} onChange={(e) => setPriority(e.target.value)}>
                                    <option value="高">高</option>
                                    <option value="中">中</option>
                                    <option value="低">低</option>
                                </Select>
                            </InputIconWrapper>
                        </FormGroup>
                        <FormGroup>
                            <SubtitleLabel>期限</SubtitleLabel>
                            <InputIconWrapper>
                                <Input
                                    type="date"
                                    value={deadline}
                                    onChange={(e) => setDeadline(e.target.value)}
                                />
                            </InputIconWrapper>
                        </FormGroup>
                        <FormGroup>
                            <SubtitleLabel>ステータス</SubtitleLabel>
                            <InputIconWrapper>
                                <Select value={status} onChange={(e) => setStatus(e.target.value)}>
                                    <option value="未着手">未着手</option>
                                    <option value="進行中">進行中</option>
                                    <option value="完了">完了</option>
                                </Select>
                            </InputIconWrapper>
                        </FormGroup>
                        <FormGroup>
                            <SubtitleLabel>ステップ</SubtitleLabel>
                            <InputIconWrapper>
                                <StyledTextarea
                                    value={steps}
                                    onChange={(e) => setSteps(e.target.value)}
                                ></StyledTextarea>
                            </InputIconWrapper>
                        </FormGroup>
                        <FormGroup>
                            <SubtitleLabel>メモ</SubtitleLabel>
                            <InputIconWrapper>
                                <StyledTextarea
                                    value={memo}
                                    onChange={(e) => setMemo(e.target.value)}
                                ></StyledTextarea>
                            </InputIconWrapper>
                        </FormGroup>
                        <FormGroup>
                            <SubtitleLabel>備考</SubtitleLabel>
                            <InputIconWrapper>
                                <StyledTextarea
                                    value={remarks}
                                    onChange={(e) => setRemarks(e.target.value)}
                                ></StyledTextarea>
                            </InputIconWrapper>
                        </FormGroup>
                        {errorMessage && <ErrorMessage>{errorMessage}</ErrorMessage>}
                        <Button type="submit">タスク追加</Button>
                    </Form>
                {/* ) : (
                    <MarkdownPreview>
                        <ReactMarkdown>{`
# タイトル
- ${title}

## 目的
${purpose}

## 期限
- ${deadline}

## 優先度
${priority}

## ステータス
${status}

### ステップ
${steps}

### メモ
${memo}

### 備考
${remarks}
                        `}</ReactMarkdown>
                    </MarkdownPreview>
                )} */}
            </ContentContainer>
        </div>
    );
};

export default CreateTask;
