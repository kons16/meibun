import React, { Component } from 'react';
import { Link } from "react-router-dom";
import axios from 'axios';
import Button from '@material-ui/core/Button';
import history from "../history";

interface State {
    name?: string
    email?: string
    password?: string
}

// フォームを入力後新規登録する。
class Signup extends Component<{}, State> {
    state: State = {
        name: "",
        email: "",
        password: "",
    };

    // tokenの有無でログイン済みかどうかチェックし、ログイン済みなら / にリダイレクト.
    // tokenは有効期限が切れてないか確認する
    componentDidMount() {
        const cookies = document.cookie;
        const cookiesArray = cookies.split(';');
        let token = "";
        const sessionName = "TEST_SESSION_KEY";

        // cookieの読み込み
        for(let c of cookiesArray){
            let cArray = c.split('=');
            if( cArray[0] === sessionName){
                token = cArray[1]
            }
        }

        if(token !== "") {
            axios.get('http://localhost:8000/signup', {withCredentials: true})
                .then((response) => {
                    if (response.data.User !== null) {
                        history.push('/')
                    } else {
                        // 有効期限が切れたtoken,またはログインしていないときcookieを削除
                        document.cookie = `${response.data.name}=; max-age=0`;
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

    // 新規登録が押されたときformの情報をpostする
    handleFormSubmit = () => {
        axios.post('http://localhost:8000/signup',
            {'name': this.state.name, 'email': this.state.email, 'password': this.state.password})
            .then((response) => {
                const name = response.data.Name;
                const token = response.data.token;
                const expiresAt = response.data.expiresAt;
                document.cookie = `${name}=${token}; expires=${expiresAt}`;
                // ここをあとで修正
                history.push('/');
            })
            .catch(() => {
                console.log("submit fail");
            });
    }

    render() {
        return (
            <div>
                <div id="form">
                    <div>
                        <span className="label">名前</span>
                        <input type="text" name="name" onChange={this.onChange} />
                    </div>
                    <div>
                        <span className="label">メールアドレス</span>
                        <input type="text" name="email" onChange={this.onChange} />
                    </div>
                    <div>
                        <span className="label">パスワード</span>
                        <input type="password" name="password" onChange={this.onChange} />
                    </div>
                    <Button variant="contained" color="primary"　style={{ marginTop: 10, width: 100 }}
                            onClick={this.handleFormSubmit} >
                        新規登録
                    </Button>
                </div>
                <Link to="/">ホームへ</Link>
            </div>
        );
    }
}

export default Signup;
