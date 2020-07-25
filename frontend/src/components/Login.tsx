import React, { Component } from 'react';
import history from "../history";
import axios from 'axios';

interface State {
    email?: string
    password?: string
}

// GET /signin でログイン済みなら / にリダイレクトさせ、未ログインならログインフォームを表示させる
class Login extends Component<{}, State> {
    state: State = {
        email: "",
        password: "",
    };

    // tokenの有無でログイン済みかどうかチェックし、ログイン済みなら / にリダイレクト.
    // tokenは有効期限が切れてないか確認する
    componentDidMount() {
        const token = localStorage.getItem('meibun_token');
        if(token !== "") {
            axios.get('http://localhost:8000/check_user', {headers: {'Authorization': token}})
                .then((response) => {
                    const userData = response.data.User;
                    if (userData == null) {
                        localStorage.setItem('meibun_token', "");
                        history.push('/')
                    }
                })
                .catch(() => {
                    console.log("index fail");
                });
        }
    }

    onChange = (e: any) => {
        this.setState({
            [e.target.name]: e.target.value,
        });
    }

    // ログインボタンが呼ばれたときの関数  /signin に POSTする
    handleFormSubmit = () => {
        axios.post('http://localhost:8000/signin',
            {'email': this.state.email, 'password': this.state.password})
            .then((response) => {
                localStorage.setItem('meibun_token', response.data.token);
            })
            .catch(() => {
                console.log("index fail");
            });
    }

    render() {
        return (
            <div>
                <span className="label">メールアドレス</span>
                <input type="text" name="email" onChange={this.onChange} />
                <span className="label">パスワード</span>
                <input type="password" name="password" onChange={this.onChange} />
                <button onClick={this.handleFormSubmit}>ログイン</button>
            </div>
        );
    }
}

export default Login;
