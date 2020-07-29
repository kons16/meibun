import React, { Component } from 'react';
import { Link, RouteComponentProps } from "react-router-dom";
import axios from "axios";
import Book from './Book';

type MyPageProps = {} & RouteComponentProps<{id: string}>;

interface MyPageState {
    books: any
}

// 自分の情報を表示するマイページ
class MyPage extends Component<MyPageProps, MyPageState> {
    constructor(props: MyPageProps) {
        super(props);
        this.state = {
            books: []
        };
    }

    componentDidMount() {
        axios.get('http://localhost:8000/users/books', {params: {id: this.props.match.params.id}})
            .then((response) => {
                const books: any[] = response.data.Books;
                if(books != null){
                    this.setState({
                        books: books
                    })
                }
                console.log(this.state.books);
            })
            .catch(() => {
                console.log("books get fail");
            });
    }

    // Bookコンポーネントに各名文情報を渡いて表示
    render() {
        return (
            <div>
                マイページです。
                <Link to="/">ホームへ</Link>
                {(() => {
                    const bookItems: any = [];
                    // console.log(this.state.books);
                    this.state.books.forEach((key: any, index: number) => {
                        bookItems.push(
                            <Book
                                key={index}
                                id={index}
                                sentence={this.state.books[index].Sentence}
                                title={this.state.books[index].Title}
                                author={this.state.books[index].Author}
                                pages={this.state.books[index].Pages}
                                harts={this.state.books[index].Harts}
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
