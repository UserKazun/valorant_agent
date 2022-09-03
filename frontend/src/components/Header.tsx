import PropTypes from "prop-types";
// import styled from "styled-components"

const Header = (props: { siteTitle: string }) => (
    <header
    style={{
      background: `#BD3944`,
      marginBottom: `1.45rem`,
      color: `#FFFFFF`,
    }}
    >
    <div
      style={{
        margin: `0 auto`,
        maxWidth: 1440,
        maxHeight: 64,
        padding: `1.45rem 1.0875rem`,
      }}
    >
      <h1
      style={{
        margin: 0,
        textAlign: 'left'
        }}>
        {props.siteTitle}
        {/* <Link
          to="/"
          style={{
            color: `white`,
            textDecoration: `none`,
          }}
        >
          {siteTitle}
        </Link> */}
      </h1>
    </div>
  </header>
)

Header.propTypes = {
    siteTitle: PropTypes.string,
}

Header.defaultProps = {
    siteTitle: `VALORANT AGENT`,
}

export default Header
