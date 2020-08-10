import React, { Component } from 'react';
import { Link, RouteComponentProps } from "react-router-dom";
import axios from "axios";
import Book from './Book';

type MyPageProps = {} & RouteComponentProps<{id: string}>;

interface MyPageState {
    books: any
    myID: number
    urlID: number
    myPageFlag: boolean     // myPageFlagがtrueのときマイページ
}

// 自分の情報を表示するマイページ
class MyPage extends Component<MyPageProps, MyPageState> {
    constructor(props: MyPageProps) {
        super(props);
        this.state = {
            books: [],
            myID: 0,
            urlID: 0,
            myPageFlag: false
        }
    }

    componentDidMount() {
        // マイページに表示するログインしているユーザー情報の取得
        axios.get('http://localhost:8000/', {withCredentials: true})
            .then((response) => {
                const userData = response.data.User;
                if(userData != null){
                    this.setState({
                        myID: userData.ID,
                        urlID: parseInt(this.props.match.params.id)
                    })
                }
            })
            .catch(() => {
                console.log("index fail");
            });

        // ユーザーに基づくbookの取得
        axios.get('http://localhost:8000/users/books', {params: {id: this.props.match.params.id}})
            .then((response) => {
                const books: any[] = response.data.Books;
                if(books != null){
                    this.setState({
                        books: books,
                        myPageFlag: (books[0].UserID === this.state.myID) ? true : false
                    })
                }
                // console.log(this.state.books);
            })
            .catch(() => {
                console.log("books get fail");
            });
    }

    // Bookコンポーネントに各名文情報を渡いて表示
    render() {
        return (
            <div>
                {this.state.urlID === this.state.myID &&
                    <div>
                        マイページです<br/>
                        <Link to={{ pathname: '/post_book', state: {myID: this.state.myID}}}>名文を追加</Link>
                    </div>
                }

                <Link to="/">ホームへ</Link> <br/>
                <Link to="/my_hart">マイハート一覧</Link>

                {(() => {
                    const bookItems: any = [];
                    // bookに削除マークを表示させるため、myPageFlagをbookに渡す
                    this.state.books.forEach((key: any, index: number) => {
                        bookItems.push(
                            <Book
                                key={index}
                                id={this.state.books[index].ID}
                                sentence={this.state.books[index].Sentence}
                                title={this.state.books[index].Title}
                                author={this.state.books[index].Author}
                                pages={this.state.books[index].Pages}
                                harts={this.state.books[index].Harts}
                                myPageFlag={this.state.myPageFlag}
                            />
                        )
                    });

                    if (this.state.books) {
                        return (
                            <div>
                                {bookItems}
                            </div>
                        )
                    }
                })()}
            </div>
        );
    }
}

export default MyPage;
