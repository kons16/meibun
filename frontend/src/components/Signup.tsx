import React, { Component } from 'react';
import { Link } from "react-router-dom";
import axios from 'axios';
import Button from '@material-ui/core/Button';
import history from "../history";

interface State {
    email?: string
    password?: string
}

class Signup extends Component<{}, State> {
    state: State = {
        email: "",
        password: "",
    };

    componentDidMount() {
    }

    onChange = (e: any) => {
        this.setState({
            [e.target.name]: e.target.value,
        });
    }

    handleFormSubmit = () => {
    }

    render() {
        return (
            <div>
                <div id="form">
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
