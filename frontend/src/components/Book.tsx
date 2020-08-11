import React, { Component } from 'react';
import axios from "axios";
import Button from "@material-ui/core/Button";
import history from "../history";

interface BookProps {
    id: number
    sentence: string
    title: string
    author: string
    pages: number
    harts: number
    myPageFlag: boolean  // trueのときはbookが自分の投稿である(バツマークと編集ボタンを表示させる)
    hartFlag: boolean   // trueのときは自分がハートをした投稿である(ハートを再度クリックでハートから削除)
    bookUserID: number  // bookに紐づくユーザー
}

interface BookState {
    book: {
        id: number
        sentence: string
        title: string
        author: string
        pages: number
        harts: number
        myPageFlag: boolean
        hartFlag: boolean
        bookUserID: number
    }
}

// Book自体のコンポーネント
class Books extends Component<BookProps, BookState> {
    constructor(props: any) {
        super(props);
        this.state = {
            book: {
                id: props.id,
                sentence: props.sentence,
                title: props.title,
                author: props.author,
                pages: props.pages,
                harts: props.harts,
                myPageFlag: props.myPageFlag,
                hartFlag: props.hartFlag,
                bookUserID: props.bookUserID
            }
        };
    }

    // bookを削除する
    handleDeleteBook = () => {
        axios.post('http://localhost:8000/delete_book',
            {'bookID': this.state.book.id},
            {withCredentials: true})
            .then((response) => {
                window.location.reload();
            })
            .catch(() => {
                console.log("delete fail");
            });
    }

    // bookの内容を編集する
    handleEditBook = () => {
        history.push({
            pathname: '/edit_book',
            state: { Book: this.state.book}
        });
    }

    // bookにハートする
    handleMakeHart = () => {
        if(this.state.book.hartFlag) {
            axios.post('http://localhost:8000/remove_hart',
                {'bookID': this.state.book.id},
                {withCredentials: true})
                .then((response) => {
                    window.location.reload();
                })
                .catch(() => {
                    console.log("delete fail");
                });
        } else {
            axios.post('http://localhost:8000/make_hart',
                {'bookID': this.state.book.id},
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
                <div id="book-sentence">{this.state.book.sentence}</div>
                <div id="book-title">『{this.state.book.title}』</div>
                <div id="book-author">{this.state.book.author}</div>
                <div id="book-pages">p.{this.state.book.pages}</div>
                <div id="book-harts">
                    {this.state.book.myPageFlag
                        ? (<div>♡</div>)
                        : (<div>
                            <button onClick={this.handleMakeHart}>
                                {this.state.book.hartFlag ? <div>☓</div> : <div>♡</div>}
                            </button>
                        </div>)
                    }
                    {this.state.book.harts}
                </div>
                {this.state.book.myPageFlag &&
                    <div>
                        <Button variant="contained" color="primary"　style={{ marginTop: 10, width: 10 }}
                                onClick={this.handleDeleteBook} >
                            ☓
                        </Button>
                        <Button variant="contained" color="primary"　style={{ marginTop: 10, marginLeft: 10, width: 10 }}
                                onClick={this.handleEditBook} >
                            編集
                        </Button>
                    </div>
                }
            </div>
        );
    }
}

export default Books;
