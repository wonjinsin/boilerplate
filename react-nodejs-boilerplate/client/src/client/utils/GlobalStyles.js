import reset from "styled-reset";
import { createGlobalStyle } from "styled-components";

const globalStyles = createGlobalStyle`
    ${reset}
    *{margin:0; padding:0; box-sizing:border-box}
    p, span, h3, h4, h5, li{word-break:keep-all;}
    body, input, button, select, textarea, sub, a {font-family:'Noto Sans KR', "맑은 고딕", Malgun Gothic, "돋움", dotum, Helvetica, sans-serif !important; color:#777; font-smoothing: antialiased; -webkit-font-smoothing: antialiased; -moz-osx-font-smoothing: grayscale;}
    img, input, button, select, textarea {vertical-align:middle;}
    ul, ol { list-style:none;}
    a { color:#333333;}
    a:hover, a:focus, a:active, a:link, a:visited { text-decoration:none;}
`;

export default globalStyles;