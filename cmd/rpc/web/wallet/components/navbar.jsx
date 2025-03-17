import Container from "react-bootstrap/Container";
import Navbar from "react-bootstrap/Navbar";
import { Nav, NavDropdown } from "react-bootstrap";
import { withTooltip } from "@/components/util";
import DarkModeToggle from "@/components/color_mode";
import { WalletLogoIcon, KeyIcon } from "@/components/svg_icons";

const navbarIconsAndTip = [
  {
    icon: (
      <svg className="navbar-image-link" id="svg" width="400" height="400" viewBox="0, 0, 400,400">
        <g id="svgg">
          <path
            id="path0"
            d="M100.284 55.006 C 65.751 72.615,54.782 152.812,80.083 202.702 C 101.517 244.965,97.692 254.946,51.362 277.651 C 12.782 296.557,-1.149 312.602,0.811 335.873 L 2.000 350.000 84.646 351.080 C 165.854 352.141,167.256 352.019,165.180 344.080 C 163.159 336.353,159.939 336.000,91.533 336.000 C 32.333 336.000,20.000 335.019,20.000 330.308 C 20.000 320.241,41.976 302.269,66.320 292.426 C 104.747 276.889,112.904 266.943,110.958 238.000 C 109.357 214.194,101.496 188.513,94.208 183.274 C 88.555 179.211,87.230 103.975,92.580 90.811 C 108.992 50.427,192.873 67.237,192.738 110.883 C 192.621 148.764,189.658 180.536,186.000 183.146 C 153.418 206.393,169.562 294.265,203.808 280.079 C 216.771 274.710,220.105 268.000,209.809 268.000 C 185.530 268.000,180.436 242.552,197.700 207.507 C 211.483 179.530,212.742 171.506,213.024 109.863 C 213.241 62.593,147.827 30.764,100.284 55.006 M251.526 118.012 C 226.527 131.598,218.376 203.987,238.033 237.848 C 252.324 262.465,248.360 273.063,220.312 285.228 C 190.603 298.112,175.154 315.977,176.792 335.555 L 178.000 350.000 288.000 350.000 L 398.000 350.000 399.218 335.388 C 400.848 315.839,383.449 295.952,354.399 284.162 C 327.528 273.256,323.766 262.306,338.295 237.288 C 350.438 216.377,352.630 152.390,341.736 136.836 C 326.363 114.889,276.442 104.472,251.526 118.012 M314.000 139.638 C 328.399 148.102,333.139 206.528,320.614 221.150 C 312.740 230.341,309.578 272.001,315.863 283.744 C 318.472 288.618,332.033 297.734,346.000 304.000 C 367.638 313.708,380.000 324.520,380.000 333.737 C 380.000 334.982,338.600 336.000,288.000 336.000 C 237.400 336.000,196.000 334.982,196.000 333.737 C 196.000 324.520,208.362 313.708,230.000 304.000 C 263.915 288.784,277.007 247.312,255.918 221.902 C 244.906 208.633,246.162 153.475,257.739 141.897 C 269.785 129.851,295.587 128.815,314.000 139.638 "
            stroke="none"
            fillRule="evenodd"
          ></path>
        </g>
      </svg>
    ),
    tip: "accounts",
  },
  {
    icon: (
      <svg className="navbar-image-link" id="svg" width="400" height="400" viewBox="0, 0, 400,400">
        <g id="svgg">
          <path
            id="path0"
            d="M108.000 59.960 L 26.000 103.437 24.791 123.718 C 23.707 141.898,24.433 144.000,31.791 144.000 C 38.240 144.000,40.000 146.565,40.000 155.966 C 40.000 163.600,42.896 169.741,48.000 172.928 C 55.511 177.619,56.000 182.256,56.000 248.743 L 56.000 319.562 37.000 320.781 L 18.000 322.000 18.000 348.000 L 18.000 374.000 200.000 374.000 L 382.000 374.000 382.000 348.000 L 382.000 322.000 363.000 320.781 L 344.000 319.562 344.000 248.743 C 344.000 182.256,344.489 177.619,352.000 172.928 C 357.104 169.741,360.000 163.600,360.000 155.966 C 360.000 146.565,361.760 144.000,368.209 144.000 C 375.567 144.000,376.293 141.898,375.209 123.719 L 374.000 103.438 290.958 59.719 C 245.285 35.674,203.885 16.109,198.958 16.242 C 194.031 16.374,153.100 36.048,108.000 59.960 M279.786 77.500 C 322.568 100.167,358.268 120.803,359.119 123.357 C 360.277 126.830,320.173 128.000,200.000 128.000 C 79.827 128.000,39.723 126.830,40.881 123.357 C 42.188 119.435,194.141 36.287,200.000 36.287 C 201.100 36.287,237.004 54.833,279.786 77.500 M120.000 145.673 C 120.000 157.192,113.644 160.000,87.572 160.000 C 72.407 160.000,60.000 159.247,60.000 158.327 C 60.000 157.407,58.907 153.807,57.572 150.327 C 55.481 144.878,59.645 144.000,87.572 144.000 C 105.407 144.000,120.000 144.753,120.000 145.673 M152.000 155.966 C 152.000 163.600,154.896 169.741,160.000 172.928 C 167.512 177.620,168.000 182.256,168.000 248.962 L 168.000 320.000 144.000 320.000 L 120.000 320.000 120.000 248.962 C 120.000 182.256,120.488 177.620,128.000 172.928 C 133.104 169.741,136.000 163.600,136.000 155.966 C 136.000 146.667,137.783 144.000,144.000 144.000 C 150.217 144.000,152.000 146.667,152.000 155.966 M232.000 145.673 C 232.000 157.192,225.644 160.000,199.572 160.000 C 184.407 160.000,172.000 159.247,172.000 158.327 C 172.000 157.407,170.907 153.807,169.572 150.327 C 167.481 144.878,171.645 144.000,199.572 144.000 C 217.407 144.000,232.000 144.753,232.000 145.673 M264.000 155.966 C 264.000 163.600,266.896 169.741,272.000 172.928 C 279.512 177.620,280.000 182.256,280.000 248.962 L 280.000 320.000 256.000 320.000 L 232.000 320.000 232.000 248.962 C 232.000 182.256,232.488 177.620,240.000 172.928 C 245.104 169.741,248.000 163.600,248.000 155.966 C 248.000 146.667,249.783 144.000,256.000 144.000 C 262.217 144.000,264.000 146.667,264.000 155.966 M344.000 145.673 C 344.000 157.192,337.644 160.000,311.572 160.000 C 296.407 160.000,284.000 159.247,284.000 158.327 C 284.000 157.407,282.907 153.807,281.572 150.327 C 279.481 144.878,283.645 144.000,311.572 144.000 C 329.407 144.000,344.000 144.753,344.000 145.673 M104.000 248.000 L 104.000 320.000 88.000 320.000 L 72.000 320.000 72.000 248.000 L 72.000 176.000 88.000 176.000 L 104.000 176.000 104.000 248.000 M216.000 248.000 L 216.000 320.000 200.000 320.000 L 184.000 320.000 184.000 248.000 L 184.000 176.000 200.000 176.000 L 216.000 176.000 216.000 248.000 M328.000 248.000 L 328.000 320.000 312.000 320.000 L 296.000 320.000 296.000 248.000 L 296.000 176.000 312.000 176.000 L 328.000 176.000 328.000 248.000 M364.000 348.044 L 364.000 360.089 199.000 359.044 L 34.000 358.000 32.734 347.000 L 31.469 336.000 197.734 336.000 L 364.000 336.000 364.000 348.044 "
            stroke="none"
            fillRule="evenodd"
          ></path>
        </g>
      </svg>
    ),
    tip: "governance",
  },
  {
    icon: (
      <svg className="navbar-image-link" id="svg" width="400" height="400" viewBox="0, 0, 400,400">
        <g id="svgg">
          <path
            id="path0"
            d="M8.000 200.000 L 8.000 344.000 200.000 344.000 L 392.000 344.000 392.000 200.000 L 392.000 56.000 200.000 56.000 L 8.000 56.000 8.000 200.000 M376.000 200.000 L 376.000 328.000 200.000 328.000 L 24.000 328.000 24.000 200.000 L 24.000 72.000 200.000 72.000 L 376.000 72.000 376.000 200.000 M120.000 95.005 C 120.000 100.737,116.508 103.698,107.663 105.467 C 52.689 116.462,36.650 199.681,82.732 234.829 C 121.252 264.210,189.926 242.443,198.222 198.223 C 200.230 187.521,202.883 184.000,208.940 184.000 C 232.078 184.000,210.161 126.260,177.596 101.421 C 158.475 86.836,120.000 82.550,120.000 95.005 M168.847 120.603 C 179.446 128.150,196.000 154.510,196.000 163.841 C 196.000 166.350,184.100 168.000,166.000 168.000 L 136.000 168.000 136.000 137.566 L 136.000 107.133 146.756 109.832 C 152.672 111.317,162.613 116.164,168.847 120.603 M256.000 128.000 C 256.000 135.467,258.667 136.000,296.000 136.000 C 333.333 136.000,336.000 135.467,336.000 128.000 C 336.000 120.533,333.333 120.000,296.000 120.000 C 258.667 120.000,256.000 120.533,256.000 128.000 M120.000 154.000 L 120.000 184.000 150.000 184.000 C 185.953 184.000,188.308 188.633,164.471 212.471 C 130.420 246.521,76.000 224.088,76.000 176.000 C 76.000 157.008,77.445 153.614,91.529 139.529 C 115.367 115.692,120.000 118.047,120.000 154.000 M256.000 176.000 C 256.000 183.467,258.667 184.000,296.000 184.000 C 333.333 184.000,336.000 183.467,336.000 176.000 C 336.000 168.533,333.333 168.000,296.000 168.000 C 258.667 168.000,256.000 168.533,256.000 176.000 M256.000 224.000 C 256.000 231.467,258.667 232.000,296.000 232.000 C 333.333 232.000,336.000 231.467,336.000 224.000 C 336.000 216.533,333.333 216.000,296.000 216.000 C 258.667 216.000,256.000 216.533,256.000 224.000 M184.000 274.000 C 184.000 279.067,174.667 280.000,124.000 280.000 C 66.667 280.000,64.000 280.356,64.000 288.000 C 64.000 295.644,66.667 296.000,124.000 296.000 C 174.667 296.000,184.000 296.933,184.000 302.000 C 184.000 305.333,187.556 308.000,192.000 308.000 C 198.933 308.000,200.000 305.333,200.000 288.000 C 200.000 270.667,198.933 268.000,192.000 268.000 C 187.556 268.000,184.000 270.667,184.000 274.000 M216.000 288.000 C 216.000 295.644,218.667 296.000,276.000 296.000 C 333.333 296.000,336.000 295.644,336.000 288.000 C 336.000 280.356,333.333 280.000,276.000 280.000 C 218.667 280.000,216.000 280.356,216.000 288.000 "
            stroke="none"
            fillRule="evenodd"
          ></path>
        </g>
      </svg>
    ),
    tip: "monitor",
  },
];

export default function Navigation({ keystore, setActiveKey, keyIdx, setNavIdx, height }) {
  return (
    <Navbar sticky="top" data-bs-theme="light" id="nav-bar">
      <Container id="nav-bar-container">
        <Navbar.Brand id="nav-bar-brand">
          <WalletLogoIcon className="nav-bar-logo" />
        </Navbar.Brand>
        <div id="nav-dropdown-container">
          <NavDropdown
            id="nav-dropdown"
            title={
              <>
                {Object.keys(keystore)[keyIdx]}
                <KeyIcon />
              </>
            }
          >
            {Object.keys(keystore).map((key, i) => (
              <NavDropdown.Item onClick={() => setActiveKey(i)} key={i} className="nav-dropdown-item">
                {key}
              </NavDropdown.Item>
            ))}
          </NavDropdown>
        </div>
        <div id="nav-bar-select-container">
          <Nav id="nav-bar-select" justify variant="tabs" defaultActiveKey="0">
            {navbarIconsAndTip.map((v, i) => (
              <Nav.Item key={i} onClick={() => setNavIdx(i)}>
                {withTooltip(<Nav.Link eventKey={i.toString()}>{v.icon}</Nav.Link>, v.tip, i, "bottom")}
              </Nav.Item>
            ))}
          </Nav>
        </div>
        <DarkModeToggle />
        <Navbar.Text className="navbar ms-4">
          Latest height :{" "}
          <strong style={{ paddingLeft: "4px" }} className="nav-bar-brand-highlight">
            {Math.max(height - 1, 1)}
          </strong>
        </Navbar.Text>
      </Container>
    </Navbar>
  );
}
