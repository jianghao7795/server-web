declare namespace Comment {
  export interface comment {
    ID: number;
    createdAt: string;
    updatedAt: string;
    articleId: number;
    article: API.Article;
    parentId: number;
    content: string;
    children: comment[];
  }
}
