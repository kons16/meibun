import React, { Component } from 'react';
import axios from "axios";
import Button from "@material-ui/core/Button";

interface BookProps {
    id: number
    sentence: string
    title: string
    author: string
    pages: number
    harts: number
    myPageFlag: boolean  // trueのときはbookが自分の投稿である(バツマークを表示させる)
    hartFlag: boolean   // trueのときは自分がハートをした投稿である(ハートを再度クリックでハートから削除)
}

interface BookState {
    id: number
    sentence: string
    title: string
    author: string
    pages: number
    harts: number
    myPageFlag: boolean
    hartFlag: boolean
}

// Book自体のコンポーネント
class Books extends Component<BookProps, BookState> {
    constructor(props: any) {
        super(props);
        this.state = {
            id: props.id,
            sentence: props.sentence,
            title: props.title,
            author: props.author,
            pages: props.pages,
            harts: props.harts,
            myPageFlag: props.myPageFlag,
            hartFlag: props.hartFlag
        };
    }

    // bookを削除する
    handleDeleteBook = () => {
        axios.post('http://localhost:8000/delete_book',
            {'bookID': this.state.id},
            {withCredentials: true})
            .then((response) => {
                window.location.reload();
            })
            .catch(() => {
                console.log("delete fail");
            });
    }

    // bookにハートする
    handleMakeHart = () => {
        if(this.state.hartFlag) {
            axios.post('http://localhost:8000/remove_hart',
                {'bookID': this.state.id},
                {withCredentials: true})
                .then((response) => {
                    window.location.reload();
                })
                .catch(() => {
                    console.log("delete fail");
                });
        } else {
            axios.post('http://localhost:8000/make_hart',
                {'bookID': this.state.id},
                {withCredentials: true})
                .then((response) => {
                    window.location.reload();
                })
                .catch(() => {
                    console.log("delete fail");
                });
        }
    }

    render() {
        return (
            <div id="book-component">
                <div id="book-sentence">{this.state.sentence}</div>
                <div id="book-title">『{this.state.title}』</div>
                <div id="book-author">{this.state.author}</div>
                <div id="book-pages">p.{this.state.pages}</div>
                <div id="book-harts">
                    {this.state.myPageFlag
                        ? (<div>♡</div>)
                        : (<div>
                            <button onClick={this.handleMakeHart}>
                                {this.state.hartFlag ? <div>☓</div> : <div>♡</div>}
                            </button>
                        </div>)
                    }
                    {this.state.harts}
                </div>
                {this.state.myPageFlag &&
                    <div>
                        <Button variant="contained" color="primary"　style={{ marginTop: 10, width: 10 }}
                                onClick={this.handleDeleteBook} >
                            ☓
                        </Button>
                    </div>
                }
            </div>
        );
    }
}

export default Books;
