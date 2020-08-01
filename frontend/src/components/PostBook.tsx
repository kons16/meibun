import React, { Component } from 'react';
import { Link } from "react-router-dom";
import history from "../history";
import axios from 'axios';
import Button from '@material-ui/core/Button';

interface State {
    sentence?: string
    title?: string
    author?: string
    pages?: number
}

interface PostBookProps {
    location: any
}

// 名文を新規登録する画面
class PostBook extends Component<PostBookProps, State> {
    constructor(props: any) {
        super(props);
        this.state = {
            sentence: "",
            title: "",
            author: "",
            pages: 0,
        };
    }

    componentDidMount() {
    }

    onChange = (e: any) => {
        this.setState({
            [e.target.name]: e.target.value,
        });
    }

    // 名文情報をpostする
    handleFormSubmit = () => {
        axios.post('http://localhost:8000/post_book',
            {'sentence': this.state.sentence, 'title': this.state.title, 'author': this.state.author, 'pages': this.state.pages},
            {withCredentials: true})
            .then((response) => {
                // マイページに遷移
                history.push({
                    pathname: `/users/${this.props.location.state.myID}`
                });
            })
            .catch(() => {
                console.log("post fail");
            });
    }

    render() {
        return (
            <div>
                <div id="form">
                    <div>
                        <span className="label">名文</span>
                        <input type="text" name="sentence" onChange={this.onChange} />
                    </div>
                    <div>
                        <span className="label">本のタイトル</span>
                        <input type="text" name="title" onChange={this.onChange} />
                    </div>
                    <div>
                        <span className="label">ページ数</span>
                        <input type="text" name="pages" onChange={this.onChange} />
                    </div>
                    <div>
                        <span className="label">著者名</span>
                        <input type="text" name="author" onChange={this.onChange} />
                    </div>
                    <Button variant="contained" color="primary"　style={{ marginTop: 10, width: 100 }}
                            onClick={this.handleFormSubmit} >
                        追加する
                    </Button>
                </div>
                <Link to="/">ホームへ</Link>
            </div>
        );
    }
}

export default PostBook;
