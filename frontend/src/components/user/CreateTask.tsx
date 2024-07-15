import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate, useParams, Navigate } from 'react-router-dom';
import Sidebar from '../common/Sidebar';
import { CreateTaskContainer, Form, FormGroup, TitleLabel, SubtitleLabel, Input, ErrorMessage, Button, InputIconWrapper, StyledTextarea, customStyles, SelectGroup } from '../../styles/CreateTaskStyles';
import { ContentContainer } from '../../styles/SidebarStyles';
import Select, { SingleValue } from 'react-select';

// オプション型の定義
interface OptionType {
    value: string;
    label: string;
}

// ユーザーが認証されているか確認する関数
const isAuthenticated = (): boolean => {
    return localStorage.getItem('authToken') !== null;
};

type PostDataType = {
    title: string | null;
    deadline: Date | null;
    priority: SingleValue<OptionType> | null;
    status: SingleValue<OptionType> | null;
    purpose: string | null;
    steps: string | null;
    memo: string | null;
    remarks: string | null;
}

const priorityOptions: OptionType[] = [
    { value: "high", label: "高" },
    { value: "middle", label: "中" },
    { value: "low", label: "低" },
];

const statusOptions: OptionType[] = [
    { value: 'yet', label: '未着手' },
    { value: 'progress', label: '進行中' },
    { value: 'done', label: '完了' },
];

const CreateTask: React.FC = () => {
    const [errorMessage, setErrorMessage] = useState('');
    const [task, setTask] = useState<PostDataType>({
        title: null,
        deadline: null,
        priority: null,
        status: null,
        purpose: null,
        steps: null,
        memo: null,
        remarks: null,
    });
    console.log(task);
    // const [title, setTitle] = useState('');
    // const [deadline, setDeadline] = useState('');
    // const [priority, setPriority] = useState<SingleValue<OptionType>>(null);
    // const [status, setStatus] = useState<SingleValue<OptionType>>(null);
    // const [purpose, setPurpose] = useState('');
    // const [steps, setSteps] = useState('');
    // const [memo, setMemo] = useState('');
    // const [remarks, setRemarks] = useState('');

    // const response = await axios.post('http://localhost:8000/api/tasks', {
    //     user_id: userId,
    //     title,
    //     deadline,
    //     priority: priority ? priority.value : '',
    //     status: status ? status.value : '',
    //     purpose,
    //     steps,
    //     memo,
    //     remarks,
    // });

    const { userId } = useParams<{ userId: string }>();

    const navigate = useNavigate();

    if (!isAuthenticated()) {
        return <Navigate to="/Signin" replace />;
    }

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setErrorMessage('');
    
        const formattedTask = {
            user_id: userId,
            title: task.title || '',
            deadline: task.deadline ? task.deadline.toISOString().split('T')[0] : '',
            priority: task.priority ? task.priority.value : '',
            status: task.status ? task.status.value : '',
            purpose: task.purpose || '',
            steps: task.steps || '',
            memo: task.memo || '',
            remarks: task.remarks || '',
        };
    
        try {
            const response = await axios.post('http://localhost:8000/api/tasks', formattedTask);
    
            console.log('登録成功:', response.data);
            const newTaskId = response.data.id;
            navigate(`/${userId}/task/${newTaskId}`);
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
                <CreateTaskContainer>
                    <Form onSubmit={handleSubmit}>
                        <FormGroup>
                            <TitleLabel>タスクタイトル</TitleLabel>
                            <InputIconWrapper>
                                {/* <Input
                                    type="text"
                                    value={title}
                                    onChange={(e) => setTitle(e.target.value)}
                                /> */}
                                <Input type="text"
                                onChange={(e) => setTask({ ...task, title: e.target.value })} />
                            </InputIconWrapper>
                        </FormGroup>
                        <FormGroup>
                            <SubtitleLabel>目的</SubtitleLabel>
                            <InputIconWrapper>
                                <Input type="text"
                                onChange={(e) => setTask({ ...task, purpose: e.target.value })} />
                            </InputIconWrapper>
                        </FormGroup>
                        <FormGroup>
                            <SubtitleLabel>優先度</SubtitleLabel>
                            <InputIconWrapper>
                                <SelectGroup>
                                    <Select
                                        // value={priority}
                                        // onChange={(option) => setPriority(option as SingleValue<OptionType>)}
                                        onChange={(value) => setTask({ ...task, priority: value as SingleValue<OptionType> })}
                                        options={priorityOptions}
                                        placeholder="優先度を選択"
                                        styles={customStyles}
                                        components={{ IndicatorSeparator: () => null }}
                                        menuPortalTarget={document.body}
                                        isMulti={false}
                                    />
                                </SelectGroup>
                            </InputIconWrapper>
                        </FormGroup>
                        <FormGroup>
                            <SubtitleLabel>期限</SubtitleLabel>
                            <InputIconWrapper>
                                <Input
                                    type="date"
                                    onChange={(e) => setTask({ ...task, deadline: new Date(e.target.value) })}
                                />
                            </InputIconWrapper>
                        </FormGroup>
                        <FormGroup>
                            <SubtitleLabel>ステータス</SubtitleLabel>
                            <InputIconWrapper>
                                <SelectGroup>
                                    <Select
                                        onChange={(value) => setTask({ ...task, status: value as SingleValue<OptionType> })}
                                        options={statusOptions}
                                        placeholder="ステータスを選択"
                                        menuPortalTarget={document.body}
                                        styles={customStyles}
                                        components={{ IndicatorSeparator: () => null }}
                                        isMulti={false}
                                    />
                                </SelectGroup>
                            </InputIconWrapper>
                        </FormGroup>
                        <FormGroup>
                            <SubtitleLabel>ステップ</SubtitleLabel>
                            <InputIconWrapper>
                                <StyledTextarea
                                    
                                />
                            </InputIconWrapper>
                        </FormGroup>
                        <FormGroup>
                            <SubtitleLabel>メモ</SubtitleLabel>
                            <InputIconWrapper>
                                <StyledTextarea
                                    onChange={(e) => setTask({ ...task, memo: e.target.value })}
                                />
                            </InputIconWrapper>
                        </FormGroup>
                        <FormGroup>
                            <SubtitleLabel>備考</SubtitleLabel>
                            <InputIconWrapper>
                                <StyledTextarea
                                    onChange={(e) => setTask({ ...task, remarks: e.target.value })}
                                />
                            </InputIconWrapper>
                        </FormGroup>
                        {errorMessage && <ErrorMessage>{errorMessage}</ErrorMessage>}
                        <Button type="submit">タスク追加</Button>
                    </Form>
                </CreateTaskContainer>
            </ContentContainer>
        </div>
    );
};

export default CreateTask;
