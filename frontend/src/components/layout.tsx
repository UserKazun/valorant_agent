import PropTypes from "prop-types"
import Header from "./header"
import "./layout.css"

type Props = {
    children : React.ReactNode;
}

const Layout: React.FC<Props> = ({ children }) => {
    return (
        <>
        <Header />
        <main>{ children }</main>
        </>
    );
};

export default Layout

Layout.propTypes = {
    children: PropTypes.node.isRequired
}
