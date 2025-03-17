// SVG map so that components are reusable

const LockIcon = () => {
  return (
    <svg xmlns="http://www.w3.org/2000/svg" width="80" height="80" className="bi bi-check-circle" viewBox="0 0 16 16">
      <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14m0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16" />
      <path d="m10.97 4.97-.02.022-3.473 4.425-2.093-2.094a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05" />
    </svg>
  );
};

const CloseIcon = () => {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="80"
      height="80"
      fill="currentColor"
      className="bi bi-c-circle"
      viewBox="0 0 16 16"
    >
      <path d="M1 8a7 7 0 1 0 14 0A7 7 0 0 0 1 8m15 0A8 8 0 1 1 0 8a8 8 0 0 1 16 0M8.146 4.992c-1.212 0-1.927.92-1.927 2.502v1.06c0 1.571.703 2.462 1.927 2.462.979 0 1.641-.586 1.729-1.418h1.295v.093c-.1 1.448-1.354 2.467-3.03 2.467-2.091 0-3.269-1.336-3.269-3.603V7.482c0-2.261 1.201-3.638 3.27-3.638 1.681 0 2.935 1.054 3.029 2.572v.088H9.875c-.088-.879-.768-1.512-1.729-1.512" />
    </svg>
  );
};

const CopyIcon = () => {
  return (
    <svg
      version="1.0"
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 100.000000 100.000000"
      preserveAspectRatio="xMidYMid meet"
    >
      <g transform="translate(0.000000,48.000000) scale(0.100000,-0.100000)" stroke="none">
        <path d="M52 428 c-8 -8 -12 -58 -12 -160 0 -141 1 -148 20 -148 19 0 20 7 20 140 l0 140 140 0 c133 0 140 1 140 20 0 19 -7 20 -148 20 -102 0 -152 -4 -160 -12z" />
        <path d="M132 348 c-17 -17 -17 -279 0 -296 17 -17 279 -17 296 0 17 17 17 279 0 296 -17 17 -279 17 -296 0z m268 -148 l0 -120 -120 0 -120 0 0 120 0 120 120 0 120 0 0 -120z" />
      </g>
    </svg>
  );
};

const DeleteOrderIcon = () => {
  return (
    <svg xmlns="http://www.w3.org/2000/svg" width="80" height="80" className="bi bi-dash-circle" viewBox="0 0 16 16">
      <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14m0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16" />
      <path d="M4 8a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7A.5.5 0 0 1 4 8" />
    </svg>
  );
};

const DiscordIcon = () => {
  return (
    <svg
      version="1.0"
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 100.000000 100.000000"
      preserveAspectRatio="xMidYMid meet"
    >
      <g transform="translate(0.000000,100.000000) scale(0.100000,-0.100000)" stroke="none">
        <path d="M300 846 c-72 -20 -129 -54 -150 -89 -47 -76 -95 -242 -106 -368 -6 -74 -6 -76 25 -111 35 -40 124 -83 191 -93 42 -6 47 -4 67 24 l23 30 -38 11 c-55 18 -122 59 -122 75 0 21 14 19 84 -14 145 -68 307 -68 452 0 70 33 84 35 84 14 0 -16 -67 -57 -122 -75 l-38 -11 23 -30 c20 -28 25 -30 67 -24 67 10 156 53 191 93 31 35 31 37 25 111 -11 126 -59 292 -106 368 -39 64 -250 135 -250 84 0 -5 15 -13 33 -16 44 -10 133 -50 142 -65 16 -26 -13 -32 -53 -11 -55 28 -154 51 -222 51 -68 0 -167 -23 -222 -51 -30 -15 -43 -18 -52 -9 -17 18 24 46 105 72 101 32 73 63 -31 34z m115 -311 c16 -15 25 -36 25 -55 0 -36 -38 -80 -70 -80 -32 0 -70 44 -70 80 0 36 38 80 70 80 12 0 32 -11 45 -25z m260 0 c16 -15 25 -36 25 -55 0 -36 -38 -80 -70 -80 -32 0 -70 44 -70 80 0 36 38 80 70 80 12 0 32 -11 45 -25z" />
      </g>
    </svg>
  );
};

const EditOrderIcon = () => {
  return (
    <svg width="80" height="80" viewBox="0 0 16 16" xmlns="http://www.w3.org/2000/svg">
      <path
        fill-rule="evenodd"
        clip-rule="evenodd"
        d="M4 4.82272C4 4.36834 4.36834 4 4.82272 4H7.33793C7.55613 4 7.76539 4.08668 7.91968 4.24097L11.759 8.08032C12.0803 8.40161 12.0803 8.92253 11.759 9.24382L9.24382 11.759C8.92253 12.0803 8.40161 12.0803 8.08032 11.759L4.24097 7.91968C4.08668 7.76539 4 7.55613 4 7.33793V4.82272ZM4.82272 4.54848C4.67126 4.54848 4.54848 4.67126 4.54848 4.82272V7.33793C4.54848 7.41066 4.57737 7.48041 4.6288 7.53184L8.46816 11.3712C8.57525 11.4783 8.74889 11.4783 8.85599 11.3712L11.3712 8.85599C11.4783 8.74889 11.4783 8.57525 11.3712 8.46816L7.53184 4.6288C7.48041 4.57737 7.41066 4.54848 7.33793 4.54848H4.82272Z"
      />
      <path
        fill-rule="evenodd"
        clip-rule="evenodd"
        d="M5.09696 6.19392C5.09696 5.58808 5.58808 5.09696 6.19392 5.09696C6.79975 5.09696 7.29087 5.58808 7.29087 6.19392C7.29087 6.79975 6.79975 7.29087 6.19392 7.29087C5.58808 7.29087 5.09696 6.79975 5.09696 6.19392ZM6.19392 5.64544C5.891 5.64544 5.64544 5.891 5.64544 6.19392C5.64544 6.49683 5.891 6.7424 6.19392 6.7424C6.49683 6.7424 6.7424 6.49683 6.7424 6.19392C6.7424 5.891 6.49683 5.64544 6.19392 5.64544Z"
      />
      <path d="M7.01664 6.19392C7.01664 6.64829 6.64829 7.01664 6.19392 7.01664C5.73954 7.01664 5.3712 6.64829 5.3712 6.19392C5.3712 5.73954 5.73954 5.3712 6.19392 5.3712C6.64829 5.3712 7.01664 5.73954 7.01664 6.19392ZM6.46816 6.19392C6.46816 6.04246 6.34538 5.91968 6.19392 5.91968C6.04246 5.91968 5.91968 6.04246 5.91968 6.19392C5.91968 6.34538 6.04246 6.46816 6.19392 6.46816C6.34538 6.46816 6.46816 6.34538 6.46816 6.19392Z" />
      <path d="M4.82272 4.27424H7.33793C7.48339 4.27424 7.6229 4.33203 7.72576 4.43489L11.5651 8.27424C11.7793 8.48843 11.7793 8.83571 11.5651 9.04991L9.04991 11.5651C8.83571 11.7793 8.48843 11.7793 8.27424 11.5651L4.43488 7.72576C4.33203 7.6229 4.27424 7.48339 4.27424 7.33793V4.82272C4.27424 4.5198 4.5198 4.27424 4.82272 4.27424ZM4.82272 7.33793L8.66207 11.1773L11.1773 8.66207L7.33793 4.82272L4.82272 4.82272L4.82272 7.33793Z" />
      <path d="M8 15C4.13401 15 1 11.866 1 8C1 4.13401 4.13401 1 8 1C11.866 1 15 4.13401 15 8C15 11.866 11.866 15 8 15ZM8 16C12.4183 16 16 12.4183 16 8C16 3.58172 12.4183 0 8 0C3.58172 0 0 3.58172 0 8C0 12.4183 3.58172 16 8 16Z" />
    </svg>
  );
};

const EditStakeIcon = () => {
  return (
    <svg xmlns="http://www.w3.org/2000/svg" width="80" height="80" className="bi bi-plus-circle" viewBox="0 0 16 16">
      <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14m0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16" />
      <path d="M8 4a.5.5 0 0 1 .5.5v3h3a.5.5 0 0 1 0 1h-3v3a.5.5 0 0 1-1 0v-3h-3a.5.5 0 0 1 0-1h3v-3A.5.5 0 0 1 8 4" />
    </svg>
  );
};

const KeyIcon = () => {
  return (
    <svg
      version="1.0"
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 100.000000 100.000000"
      preserveAspectRatio="xMidYMid meet"
    >
      <g transform="translate(0.000000,100.000000) scale(0.100000,-0.100000)" stroke="none">
        <path d="M604 991 c-52 -13 -135 -64 -165 -100 -65 -79 -93 -193 -69 -280 l12 -44 -191 -191 -191 -191 0 -66 c0 -58 3 -68 28 -92 25 -24 36 -27 93 -27 62 0 67 2 102 38 20 21 37 48 37 60 0 16 6 22 23 22 24 0 97 66 97 88 0 7 6 12 13 12 16 0 67 52 67 68 0 7 5 12 12 12 7 0 31 18 54 41 41 40 41 40 85 29 87 -24 206 7 286 74 16 14 45 53 64 88 31 58 34 70 33 148 0 72 -4 93 -26 136 -39 75 -70 107 -137 143 -65 34 -164 49 -227 32z m166 -47 c105 -31 190 -149 190 -264 0 -150 -130 -280 -280 -280 -150 0 -280 130 -280 280 0 149 126 277 276 279 22 1 64 -6 94 -15z m-346 -454 c14 -22 44 -52 66 -66 l41 -27 -28 -29 c-17 -18 -37 -28 -55 -28 -24 0 -28 -4 -28 -28 0 -36 -16 -52 -52 -52 -24 0 -28 -4 -28 -27 0 -41 -36 -73 -82 -73 -37 0 -38 -1 -38 -37 0 -58 -30 -83 -101 -83 -32 0 -61 5 -65 11 -4 7 65 83 165 183 136 136 169 174 159 184 -10 10 -46 -21 -175 -150 -162 -161 -163 -162 -163 -128 0 32 16 50 177 212 98 98 178 178 179 178 1 0 13 -18 28 -40z" />
        <path d="M583 875 c-28 -15 -56 -32 -62 -39 -8 -9 30 -52 149 -171 l158 -158 31 48 c53 83 54 162 4 243 -58 94 -179 127 -280 77z m173 -34 c78 -35 117 -117 95 -199 -22 -85 -22 -85 -157 49 -68 68 -120 127 -117 132 10 16 67 36 103 36 19 0 53 -8 76 -18z" />
      </g>
    </svg>
  );
};

const PauseIcon = () => {
  return (
    <svg xmlns="http://www.w3.org/2000/svg" width="80" height="80" viewBox="0 0 16 16">
      <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14m0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16" />
      <path d="M5.5 5.75a0.75 0.75 0 1 1 1.5 0v4.5a0.75 0.75 0 1 1-1.5 0zm3.75 0a0.75 0.75 0 1 1 1.5 0v4.5a0.75 0.75 0 1 1-1.5 0z" />
    </svg>
  );
};

const PauseFilledIcon = () => {
  return (
    <svg
      version="1.0"
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 100.000000 100.000000"
      preserveAspectRatio="xMidYMid meet"
    >
      <path
        d="M0 0 C10.06866931 7.25087772 17.72756886 16.79280068 20.3125 29.25 C21.70668558 45.91292143 19.18676434 58.15716957 8.640625 71.33984375 C2.22382219 78.48333577 -8.09918475 84.61508614 -17.80297852 85.42700195 C-34.86566428 85.88895149 -46.79126814 83.44661518 -59.8125 71.75 C-70.07212814 61.14084545 -72.17133125 48.37416494 -71.96484375 34.2578125 C-71.50503107 21.94581333 -64.87448062 11.59167312 -56.171875 3.17578125 C-39.79726859 -9.51711334 -17.57628553 -10.86641977 0 0 Z M-41.6875 21.25 C-41.6875 33.13 -41.6875 45.01 -41.6875 57.25 C-38.3875 57.25 -35.0875 57.25 -31.6875 57.25 C-31.6875 45.37 -31.6875 33.49 -31.6875 21.25 C-34.9875 21.25 -38.2875 21.25 -41.6875 21.25 Z M-19.6875 21.25 C-19.6875 33.13 -19.6875 45.01 -19.6875 57.25 C-16.3875 57.25 -13.0875 57.25 -9.6875 57.25 C-9.6875 45.37 -9.6875 33.49 -9.6875 21.25 C-12.9875 21.25 -16.2875 21.25 -19.6875 21.25 Z "
        transform="translate(75.6875,10.75)"
      />
    </svg>
  );
};

const PollIcon = () => {
  return (
    <svg
      version="1.0"
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 100.000000 100.000000"
      preserveAspectRatio="xMidYMid meet"
    >
      <path
        d="M0 0 C8.93401633 6.55707621 14.66257803 13.6752341 18.1875 24.25 C19.34652003 43.28832896 19.34652003 43.28832896 14.1875 51.25 C13.589375 52.28125 12.99125 53.3125 12.375 54.375 C5.09736106 63.93989689 -3.80266298 70.38166134 -15.70605469 72.96435547 C-20.83999955 73.6605441 -26.14852082 73.30820028 -31.30151367 72.92700195 C-37.44888769 72.63702602 -40.40145783 73.80811534 -45.08203125 77.78613281 C-49.15565897 81.23216439 -52.94282901 81.48889387 -58.125 81.375 C-59.19363281 81.35695313 -60.26226562 81.33890625 -61.36328125 81.3203125 C-62.17152344 81.29710938 -62.97976562 81.27390625 -63.8125 81.25 C-63.379375 80.35796875 -62.94625 79.4659375 -62.5 78.546875 C-61.6646875 76.79117188 -61.6646875 76.79117188 -60.8125 75 C-60.255625 73.83984375 -59.69875 72.6796875 -59.125 71.484375 C-57.56182213 68.21401402 -57.56182213 68.21401402 -57.8125 64.25 C-59.36123579 62.58901903 -59.36123579 62.58901903 -61.4375 61.125 C-67.41029095 56.21121454 -73.12866887 48.23847214 -74.03173828 40.33007812 C-74.71038309 27.80212458 -73.85579691 18.62791645 -65.375 8.875 C-61.06226969 4.26975407 -56.66117869 0.46016219 -51 -2.375 C-50.20980469 -2.77976562 -49.41960938 -3.18453125 -48.60546875 -3.6015625 C-33.42144948 -9.84506133 -13.93272783 -8.64301421 0 0 Z M-62.75 13.6875 C-68.79888252 22.19017449 -69.87982814 29.80457517 -68.8125 40.25 C-66.63868991 48.49082809 -60.89332664 56.5294489 -53.8125 61.25 C-53.1525 61.25 -52.4925 61.25 -51.8125 61.25 C-52.4725 65.87 -53.1325 70.49 -53.8125 75.25 C-48.77773198 74.34594413 -48.77773198 74.34594413 -45.69921875 71.1953125 C-41.81014865 67.67507882 -38.7245558 67.71105075 -33.62768555 67.92749023 C-32.63679932 67.99266846 -31.64591309 68.05784668 -30.625 68.125 C-17.19155543 68.5483916 -7.15500252 66.33995031 3.1875 57.25 C10.15645642 49.93917025 13.27656911 43.44683044 13.625 33.375 C13.35806345 23.3702182 10.35324558 16.52141111 3.453125 9.2890625 C-7.05029227 -0.16401304 -18.68622464 -2.53871016 -32.3984375 -2.09375 C-44.57616079 -1.36485677 -55.01912833 4.30694441 -62.75 13.6875 Z "
        transform="translate(77.8125,14.75)"
      />
      <path
        d="M0 0 C3.96 0 7.92 0 12 0 C12 13.2 12 26.4 12 40 C8.04 40 4.08 40 0 40 C0 26.8 0 13.6 0 0 Z M4 4 C4 14.56 4 25.12 4 36 C5.32 36 6.64 36 8 36 C8 25.44 8 14.88 8 4 C6.68 4 5.36 4 4 4 Z "
        transform="translate(60,28)"
      />
      <path
        d="M0 0 C3.96 0 7.92 0 12 0 C12 10.56 12 21.12 12 32 C8.04 32 4.08 32 0 32 C0 21.44 0 10.88 0 0 Z M4 4 C4 11.92 4 19.84 4 28 C5.32 28 6.64 28 8 28 C8 20.08 8 12.16 8 4 C6.68 4 5.36 4 4 4 Z "
        transform="translate(28,36)"
      />
      <path
        d="M0 0 C3.96 0 7.92 0 12 0 C12 7.26 12 14.52 12 22 C8.04 22 4.08 22 0 22 C0 14.74 0 7.48 0 0 Z M4 4 C4 8.62 4 13.24 4 18 C5.32 18 6.64 18 8 18 C8 13.38 8 8.76 8 4 C6.68 4 5.36 4 4 4 Z "
        transform="translate(44,46)"
      />
    </svg>
  );
};

const ProposalIcon = () => {
  return (
    <svg
      version="1.0"
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 100.000000 100.000000"
      preserveAspectRatio="xMidYMid meet"
    >
      <g transform="translate(0.000000,100.000000) scale(0.100000,-0.100000)" stroke="none">
        <path d="M140 500 l0 -460 360 0 360 0 0 460 0 460 -360 0 -360 0 0 -460z m380 296 c0 -77 4 -127 11 -131 5 -4 28 5 50 20 l39 28 39 -28 c22 -15 45 -24 50 -20 7 4 11 54 11 131 l0 124 50 0 50 0 0 -420 0 -420 -320 0 -320 0 0 420 0 420 170 0 170 0 0 -124z m160 24 l0 -101 -30 23 -30 22 -30 -22 -30 -23 0 101 0 100 60 0 60 0 0 -100z" />
        <path d="M285 529 c-16 -25 19 -29 221 -27 178 3 209 5 209 18 0 13 -31 15 -212 18 -147 2 -214 -1 -218 -9z" />
        <path d="M285 429 c-16 -24 17 -29 171 -27 133 3 159 5 159 18 0 13 -26 15 -162 18 -110 2 -164 -1 -168 -9z" />
        <path d="M285 329 c-16 -25 19 -29 221 -27 178 3 209 5 209 18 0 13 -31 15 -212 18 -147 2 -214 -1 -218 -9z" />
        <path d="M285 229 c-16 -24 17 -29 171 -27 133 3 159 5 159 18 0 13 -26 15 -162 18 -110 2 -164 -1 -168 -9z" />
      </g>
    </svg>
  );
};

const SearchIcon = () => {
  return (
    <svg
      version="1.0"
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 100.000000 100.000000"
      preserveAspectRatio="xMidYMid meet"
    >
      <g transform="translate(0.000000,100.000000) scale(0.100000,-0.100000)" stroke="none">
        <path
          d="M0 0 C0.64324219 0.350625 1.28648438 0.70125 1.94921875 1.0625 C7.09252599 4.27567242 10.07007011 8.77939764 13 14 C13.34289062 14.59296875 13.68578125 15.1859375 14.0390625 15.796875 C17.49225073 23.09052491 17.23638693 33.4201442 15.16796875 41.08203125 C14.21555527 43.4612814 13.15718358 45.71519593 12 48 C11.34 49.65 10.68 51.3 10 53 C17.92 60.92 25.84 68.84 34 77 C32.38262276 80.23475448 31.52370777 81.70572021 29 84 C25.13084637 82.42824389 22.82112822 79.96502822 19.97265625 76.95703125 C19.0064484 75.95581145 18.03963307 74.95517759 17.07226562 73.95507812 C15.56026972 72.38101625 14.05343367 70.80287502 12.55688477 69.21411133 C11.10471536 67.67737259 9.63436768 66.16009556 8.16015625 64.64453125 C7.29011963 63.73034424 6.42008301 62.81615723 5.52368164 61.87426758 C3.17316459 59.71743089 3.17316459 59.71743089 0.38452148 60.11206055 C-2.18849481 61.03209602 -2.18849481 61.03209602 -5.3125 62.5625 C-13.17470167 65.57609533 -23.75297758 65.25452529 -31.5625 62.3125 C-33.06479531 61.58777463 -34.54717643 60.81940328 -36 60 C-36.96486328 59.4740625 -36.96486328 59.4740625 -37.94921875 58.9375 C-43.09252599 55.72432758 -46.07007011 51.22060236 -49 46 C-49.34289062 45.40703125 -49.68578125 44.8140625 -50.0390625 44.203125 C-53.50668093 36.87899635 -53.13907128 26.72163418 -51.2734375 18.98828125 C-50.34436871 16.52209457 -49.29375631 14.29386398 -48 12 C-47.649375 11.35675781 -47.29875 10.71351562 -46.9375 10.05078125 C-43.72432758 4.90747401 -39.22060236 1.92992989 -34 -1 C-33.11054687 -1.51433594 -33.11054687 -1.51433594 -32.203125 -2.0390625 C-22.517475 -6.62474635 -9.06835379 -5.11461884 0 0 Z M-42 12 C-47.09105238 20.00442917 -47.71612339 27.62901394 -47 37 C-45.19183983 44.36335107 -40.28636615 51.27374506 -33.875 55.3515625 C-25.8371213 59.45915493 -17.23604399 60.53133625 -8.359375 58.3515625 C-0.39234902 55.0923246 5.79691768 49.67058816 9.30859375 41.796875 C12.11714696 34.49781017 12.4312509 25.44305775 9.4375 18.125 C4.68335744 9.16719455 -1.49614479 4.16795174 -11 1 C-23.88681747 -0.010406 -33.22519621 2.15316136 -42 12 Z "
          transform="translate(60,10)"
        />
      </g>
    </svg>
  );
};

const SendIcon = () => {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="80"
      height="80"
      className="bi bi-arrow-up-circle"
      viewBox="0 0 16 16"
    >
      <path
        fill-rule="evenodd"
        d="M1 8a7 7 0 1 0 14 0A7 7 0 0 0 1 8m15 0A8 8 0 1 1 0 8a8 8 0 0 1 16 0m-7.5 3.5a.5.5 0 0 1-1 0V5.707L5.354 7.854a.5.5 0 1 1-.708-.708l3-3a.5.5 0 0 1 .708 0l3 3a.5.5 0 0 1-.708.708L8.5 5.707z"
      />
    </svg>
  );
};

const StakeIcon = () => {
  return (
    <svg width="80" height="80" viewBox="0 0 16 16" xmlns="http://www.w3.org/2000/svg">
      <path
        d="M8 4.4375C8.44183 4.4375 8.8 4.82925 8.8 5.3125V7.0625H7.2V5.3125C7.2 4.82925 7.55817 4.4375 8 4.4375ZM9.2 7.0625V5.3125C9.2 4.58763 8.66274 4 8 4C7.33726 4 6.8 4.58763 6.8 5.3125V7.0625C6.35817 7.0625 6 7.45425 6 7.9375V10.125C6 10.6082 6.35817 11 6.8 11H9.2C9.64183 11 10 10.6082 10 10.125V7.9375C10 7.45425 9.64183 7.0625 9.2 7.0625ZM6.8 7.5H9.2C9.42091 7.5 9.6 7.69588 9.6 7.9375V10.125C9.6 10.3666 9.42091 10.5625 9.2 10.5625H6.8C6.57909 10.5625 6.4 10.3666 6.4 10.125V7.9375C6.4 7.69588 6.57909 7.5 6.8 7.5Z"
        stroke="currentColor"
        strokeWidth=".8"
      />
      <path d="M8 15C4.13401 15 1 11.866 1 8C1 4.13401 4.13401 1 8 1C11.866 1 15 4.13401 15 8C15 11.866 11.866 15 8 15ZM8 16C12.4183 16 16 12.4183 16 8C16 3.58172 12.4183 0 8 0C3.58172 0 0 3.58172 0 8C0 12.4183 3.58172 16 8 16Z" />
    </svg>
  );
};

const SwapIcon = () => {
  return (
    <svg xmlns="http://www.w3.org/2000/svg" width="80" height="80" viewBox="0 0 16 16">
      <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14m0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16" />
      <g transform="scale(0.66) translate(4 4)">
        <path d="M11.534 7h3.932a.25.25 0 0 1 .192.41l-1.966 2.36a.25.25 0 0 1-.384 0l-1.966-2.36a.25.25 0 0 1 .192-.41m-11 2h3.932a.25.25 0 0 0 .192-.41L2.692 6.23a.25.25 0 0 0-.384 0L.342 8.59A.25.25 0 0 0 .534 9" />
        <path d="M8 3c-1.552 0-2.94.707-3.857 1.818a.5.5 0 1 1-.771-.636A6.002 6.002 0 0 1 13.917 7H12.9A5 5 0 0 0 8 3M3.1 9a5.002 5.002 0 0 0 8.757 2.182.5.5 0 1 1 .771.636A6.002 6.002 0 0 1 2.083 9z" />
      </g>
    </svg>
  );
};

const TransactionIcon = () => {
  return (
    <svg
      version="1.0"
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 100.000000 100.000000"
      preserveAspectRatio="xMidYMid meet"
    >
      <g transform="translate(0.000000,100.000000) scale(0.100000,-0.100000)" stroke="none">
        <path d="M396 949 c-66 -16 -161 -68 -207 -114 l-44 -44 -5 72 c-4 55 -9 72 -20 72 -12 0 -16 -20 -18 -108 l-3 -108 108 3 c88 2 108 6 108 18 0 12 -17 16 -74 20 l-75 5 48 42 c244 214 613 103 692 -207 63 -246 -110 -490 -367 -517 -49 -5 -59 -9 -59 -24 0 -15 8 -19 38 -19 151 1 325 116 395 260 125 260 -25 579 -304 645 -71 16 -153 18 -213 4z" />
        <path d="M480 752 c0 -24 -6 -31 -40 -45 -89 -35 -108 -154 -32 -199 15 -9 56 -24 92 -33 74 -19 100 -38 100 -72 0 -37 -16 -60 -51 -72 -65 -23 -120 -7 -140 40 -7 18 -19 29 -31 29 -17 0 -19 -4 -13 -27 9 -38 41 -70 81 -83 28 -10 34 -16 34 -41 0 -22 5 -29 20 -29 15 0 20 7 20 29 0 25 6 31 34 41 55 18 81 54 81 112 0 43 -4 53 -29 73 -17 13 -61 31 -99 41 -77 19 -107 41 -107 80 0 39 50 84 95 84 38 0 85 -35 85 -65 0 -9 9 -15 21 -15 18 0 20 4 15 26 -11 41 -34 71 -66 84 -24 10 -30 19 -30 41 0 22 -5 29 -20 29 -15 0 -20 -7 -20 -28z" />
        <path d="M40 500 c0 -13 7 -20 20 -20 13 0 20 7 20 20 0 13 -7 20 -20 20 -13 0 -20 -7 -20 -20z" />
        <path d="M59 442 c-11 -2 -15 -10 -12 -23 6 -20 28 -26 38 -9 9 15 -8 36 -26 32z" />
        <path d="M67 363 c-3 -5 -2 -15 3 -22 10 -17 43 -6 38 12 -6 15 -34 22 -41 10z" />
        <path d="M107 299 c-14 -8 -15 -14 -6 -31 10 -18 13 -19 26 -5 21 21 5 50 -20 36z" />
        <path d="M145 230 c-4 -6 -1 -17 5 -26 10 -11 16 -12 26 -3 11 9 11 15 3 25 -14 17 -25 18 -34 4z" />
        <path d="M202 177 c-10 -11 -9 -17 3 -27 24 -20 45 4 24 26 -14 13 -17 13 -27 1z" />
        <path d="M261 126 c-10 -12 -8 -17 8 -25 16 -9 22 -8 30 8 8 15 7 20 -5 25 -21 8 -20 8 -33 -8z" />
        <path d="M333 94 c-7 -18 17 -38 31 -24 11 11 1 40 -14 40 -6 0 -13 -7 -17 -16z" />
        <path d="M407 83 c-13 -12 -7 -31 12 -36 24 -6 37 29 14 37 -18 7 -18 7 -26 -1z" />
      </g>
    </svg>
  );
};

const TwitterIcon = () => {
  return (
    <svg
      version="1.0"
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 100.000000 100.000000"
      preserveAspectRatio="xMidYMid meet"
    >
      <path
        d="M0 0 C2.078125 1.07421875 2.078125 1.07421875 3.859375 2.1640625 C5.8382613 3.28393225 7.31326273 3.86322595 9.578125 4.13671875 C11.9179429 3.39881648 11.9179429 3.39881648 14.203125 2.19921875 C18.41202331 0.13671875 18.41202331 0.13671875 20.578125 0.13671875 C19.4687178 3.62342709 18.31474834 6.91156112 16.578125 10.13671875 C18.558125 9.47671875 20.538125 8.81671875 22.578125 8.13671875 C22.10387643 12.40495589 19.62051097 14.31164606 16.578125 17.13671875 C16.03671875 17.56339844 15.4953125 17.99007813 14.9375 18.4296875 C12.98189186 20.8854368 13.24180473 22.77686885 13.203125 25.88671875 C12.25679038 42.53107469 4.44364581 56.67636206 -7.421875 68.13671875 C-20.46147688 79.47742157 -36.35637178 81.96152974 -53.00976562 81.32470703 C-61.71639985 80.64615355 -69.86189128 77.37030963 -77.421875 73.13671875 C-77.421875 72.47671875 -77.421875 71.81671875 -77.421875 71.13671875 C-76.37128906 71.03875 -75.32070312 70.94078125 -74.23828125 70.83984375 C-66.33967215 70.02628137 -59.45499737 69.09928282 -52.421875 65.13671875 C-51.431875 64.64171875 -51.431875 64.64171875 -50.421875 64.13671875 C-51.08574219 63.84539062 -51.74960937 63.5540625 -52.43359375 63.25390625 C-58.1305198 60.6092019 -61.98676033 58.45560599 -65.421875 53.13671875 C-66.234375 50.26171875 -66.234375 50.26171875 -66.421875 48.13671875 C-65.101875 48.13671875 -63.781875 48.13671875 -62.421875 48.13671875 C-62.96457031 47.70101562 -63.50726563 47.2653125 -64.06640625 46.81640625 C-69.7573635 42.07906886 -73.10432026 38.28640487 -74.109375 30.76171875 C-74.2125 29.56546875 -74.315625 28.36921875 -74.421875 27.13671875 C-72.441875 27.46671875 -70.461875 27.79671875 -68.421875 28.13671875 C-68.7828125 27.71003906 -69.14375 27.28335938 -69.515625 26.84375 C-74.11195215 20.65147592 -74.3687968 14.67222276 -73.421875 7.13671875 C-72.39453125 4.74609375 -72.39453125 4.74609375 -71.421875 3.13671875 C-68.16219665 4.5637335 -66.06437698 6.33120609 -63.546875 8.82421875 C-54.59546364 16.9084162 -41.69047306 23.13671875 -29.421875 23.13671875 C-29.48375 22.14671875 -29.545625 21.15671875 -29.609375 20.13671875 C-29.55442232 14.6854125 -28.33504108 10.20829986 -24.58203125 6.14453125 C-17.32405307 -0.37007859 -9.79377932 -3.49100843 0 0 Z "
        transform="translate(77.421875,11.86328125)"
      />
    </svg>
  );
};

const UnpauseIcon = () => {
  return (
    <svg xmlns="http://www.w3.org/2000/svg" width="80" height="80" className="bi bi-play-circle" viewBox="0 0 16 16">
      <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14m0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16" />
      <path d="M6.271 5.055a.5.5 0 0 1 .52.038l3.5 2.5a.5.5 0 0 1 0 .814l-3.5 2.5A.5.5 0 0 1 6 10.5v-5a.5.5 0 0 1 .271-.445" />
    </svg>
  );
};

const UnstakeIcon = () => {
  return (
    <svg width="80" height="80" viewBox="0 0 16 16" xmlns="http://www.w3.org/2000/svg">
      <path
        d="M9.84615 4.4375C9.42132 4.4375 9.07692 4.82925 9.07692 5.3125V7.0625C9.50176 7.0625 9.84615 7.45425 9.84615 7.9375V10.125C9.84615 10.6082 9.50176 11 9.07692 11H6.76923C6.3444 11 6 10.6082 6 10.125V7.9375C6 7.45425 6.3444 7.0625 6.76923 7.0625H8.69231V5.3125C8.69231 4.58763 9.2089 4 9.84615 4C10.4834 4 11 4.58763 11 5.3125V7.0625C11 7.18331 10.9139 7.28125 10.8077 7.28125C10.7015 7.28125 10.6154 7.18331 10.6154 7.0625V5.3125C10.6154 4.82925 10.271 4.4375 9.84615 4.4375ZM6.76923 7.5C6.55681 7.5 6.38462 7.69588 6.38462 7.9375V10.125C6.38462 10.3666 6.55681 10.5625 6.76923 10.5625H9.07692C9.28934 10.5625 9.46154 10.3666 9.46154 10.125V7.9375C9.46154 7.69588 9.28934 7.5 9.07692 7.5H6.76923Z"
        stroke="currentColor"
        strokeWidth=".8"
      />
      <path d="M8 15C4.13401 15 1 11.866 1 8C1 4.13401 4.13401 1 8 1C11.866 1 15 4.13401 15 8C15 11.866 11.866 15 8 15ZM8 16C12.4183 16 16 12.4183 16 8C16 3.58172 12.4183 0 8 0C3.58172 0 0 3.58172 0 8C0 12.4183 3.58172 16 8 16Z" />
    </svg>
  );
};

const VoteIcon = () => {
  return (
    <svg
      version="1.0"
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 100.000000 100.000000"
      preserveAspectRatio="xMidYMid meet"
    >
      <g transform="translate(0.000000,100.000000) scale(0.100000,-0.100000)" stroke="none">
        <path d="M152 900 c-18 -11 -41 -34 -52 -52 -19 -31 -20 -51 -20 -349 0 -449 -30 -421 432 -417 439 4 408 -27 408 418 0 341 -2 353 -69 397 -34 23 -38 23 -350 23 -298 0 -318 -1 -349 -20z m677 -34 c48 -26 51 -44 51 -366 0 -322 -3 -340 -51 -366 -37 -19 -613 -21 -654 -2 -52 24 -55 41 -55 370 1 324 3 337 51 364 36 20 619 20 658 0z" />
        <path d="M591 525 l-132 -155 -82 74 c-66 61 -84 73 -95 63 -12 -10 1 -26 77 -99 50 -49 96 -88 101 -88 14 0 291 331 287 343 -9 28 -33 7 -156 -138z" />
      </g>
    </svg>
  );
};

const WalletLogoIcon = () => {
  /*
     TODO: Should the logo icon be included in this library? Storing for now for ease of use, but will need to
     revisit when building the brand constructor.
     */
  return (
    <svg width="300" viewBox="0 0 4575 316" fill="none" xmlns="http://www.w3.org/2000/svg">
      <g clip-path="url(#clip0_303_5)">
        <path
          d="M218.807 14.2538V63.571L249.338 33.0401C239.772 25.4809 229.516 19.2188 218.807 14.2538Z"
          fill="black"
        />
        <path
          d="M183.395 2.81558V98.9827L147.181 135.197V0.150111C159.292 -0.383373 171.454 0.505117 183.395 2.81558Z"
          fill="black"
        />
        <path
          d="M111.769 170.608L111.769 5.87454C99.0116 9.49921 86.6301 14.8012 74.9728 21.7806L74.9728 207.405L111.769 170.608Z"
          fill="black"
        />
        <path
          d="M39.561 50.9423L39.561 242.817L33.0401 249.338C-13.0892 190.961 -10.9155 106.898 39.561 50.9423Z"
          fill="black"
        />
        <path
          d="M172.288 160.169L208.502 123.955H305.057C307.42 135.891 308.36 148.053 307.877 160.169L172.288 160.169Z"
          fill="black"
        />
        <path
          d="M243.914 88.5435H293.443C288.413 77.8272 282.081 67.571 274.444 58.0133L243.914 88.5435Z"
          fill="black"
        />
        <path
          d="M136.877 195.581L100.08 232.377H286.606C293.512 220.712 298.747 208.331 302.309 195.581H136.877Z"
          fill="black"
        />
        <path
          d="M58.0133 274.444L64.6685 267.789H257.774C201.614 319.085 116.663 321.303 58.0133 274.444Z"
          fill="black"
        />
        <path
          d="M672 187.2C669.333 211.733 660.667 233.467 646 252.4C631.6 271.067 612.667 285.6 589.2 296C565.733 306.4 538.933 311.6 508.8 311.6C475.2 311.6 445.733 305.067 420.4 292C395.067 278.933 375.333 260.8 361.2 237.6C347.067 214.133 340 186.933 340 156C340 125.067 347.067 98 361.2 74.8C375.333 51.3333 395.067 33.0667 420.4 20C445.733 6.93331 475.2 0.399977 508.8 0.399977C538.933 0.399977 565.733 5.59998 589.2 16C612.667 26.4 631.6 40.9333 646 59.6C660.667 78.2667 669.333 100 672 124.8H618.8C615.867 108.8 609.467 95.0667 599.6 83.6C590 72.1333 577.467 63.3333 562 57.2C546.533 50.8 528.667 47.6 508.4 47.6C484.933 47.6 464.533 52 447.2 60.8C430.133 69.6 416.933 82.1333 407.6 98.4C398.267 114.667 393.6 133.867 393.6 156C393.6 178.133 398.267 197.333 407.6 213.6C416.933 229.867 430.133 242.4 447.2 251.2C464.533 260 484.933 264.4 508.4 264.4C528.667 264.4 546.533 261.333 562 255.2C577.467 249.067 590 240.267 599.6 228.8C609.467 217.067 615.867 203.2 618.8 187.2H672ZM746.772 237.2V193.6H956.372V237.2H746.772ZM885.972 5.99998L1029.97 306H973.572L842.772 26H863.172L732.372 306H675.972L819.972 5.99998H885.972ZM1344.54 274L1322.94 276.4V5.99998H1374.14V306H1306.94L1096.54 35.6L1117.74 33.2V306H1066.54V5.99998H1135.34L1344.54 274ZM1597.15 311.6C1563.01 311.6 1533.01 305.067 1507.15 292C1481.28 278.933 1461.15 260.8 1446.75 237.6C1432.35 214.133 1425.15 186.933 1425.15 156C1425.15 125.067 1432.35 98 1446.75 74.8C1461.15 51.3333 1481.28 33.0667 1507.15 20C1533.01 6.93331 1563.01 0.399977 1597.15 0.399977C1631.28 0.399977 1661.28 6.93331 1687.15 20C1713.01 33.0667 1733.15 51.3333 1747.55 74.8C1761.95 98 1769.15 125.067 1769.15 156C1769.15 186.933 1761.95 214.133 1747.55 237.6C1733.15 260.8 1713.01 278.933 1687.15 292C1661.28 305.067 1631.28 311.6 1597.15 311.6ZM1597.15 264.4C1621.41 264.4 1642.35 260 1659.95 251.2C1677.55 242.4 1691.15 229.867 1700.75 213.6C1710.61 197.333 1715.55 178.133 1715.55 156C1715.55 133.867 1710.61 114.667 1700.75 98.4C1691.15 82.1333 1677.55 69.6 1659.95 60.8C1642.35 52 1621.41 47.6 1597.15 47.6C1573.15 47.6 1552.21 52 1534.35 60.8C1516.75 69.6 1503.01 82.1333 1493.15 98.4C1483.55 114.667 1478.75 133.867 1478.75 156C1478.75 178.133 1483.55 197.333 1493.15 213.6C1503.01 229.867 1516.75 242.4 1534.35 251.2C1552.21 260 1573.15 264.4 1597.15 264.4ZM1979.24 5.99998C2001.64 5.99998 2020.98 9.99998 2037.24 18C2053.78 26 2066.44 37.2 2075.24 51.6C2084.31 65.7333 2088.84 82.1333 2088.84 100.8C2088.84 119.467 2084.31 136 2075.24 150.4C2066.44 164.533 2053.78 175.6 2037.24 183.6C2020.98 191.6 2001.64 195.6 1979.24 195.6H1844.44V150H1977.24C1995.38 150 2009.51 145.6 2019.64 136.8C2030.04 128 2035.24 116 2035.24 100.8C2035.24 85.6 2030.04 73.6 2019.64 64.8C2009.51 56 1995.38 51.6 1977.24 51.6H1849.64L1873.24 25.2V306H1820.44V5.99998H1979.24ZM2231.33 306V162.8H2284.13V306H2231.33ZM2089.73 5.99998H2149.33L2271.73 159.6H2243.33L2365.73 5.99998H2425.33L2274.53 195.2L2240.93 195.6L2089.73 5.99998ZM2903.49 286H2883.49L2975.49 5.99998H3030.69L2929.09 306H2861.89L2760.69 31.2H2785.49L2684.29 306H2617.09L2515.49 5.99998H2571.09L2662.69 286L2643.49 285.6L2744.69 5.99998H2801.89L2903.49 286ZM3083.1 237.2V193.6H3292.7V237.2H3083.1ZM3222.3 5.99998L3366.3 306H3309.9L3179.1 26H3199.5L3068.7 306H3012.3L3156.3 5.99998H3222.3ZM3455.67 5.99998V282.4L3432.07 259.6H3658.07V306H3402.87V5.99998H3455.67ZM3750.98 5.99998V282.4L3727.38 259.6H3953.38V306H3698.18V5.99998H3750.98ZM4233.09 134V178H4023.49V134H4233.09ZM4058.29 156L4037.09 290L4011.89 260.4H4244.69V306H3983.09L4007.09 156L3983.09 5.99998H4242.69V51.6H4011.89L4037.09 22L4058.29 156ZM4391.07 30.4H4443.87V306H4391.07V30.4ZM4267.07 5.99998H4567.87V52.4H4267.07V5.99998Z"
          fill="black"
        />
      </g>
      <defs>
        <clipPath id="clip0_303_5">
          <rect width="4575" height="316" fill="white" />
        </clipPath>
      </defs>
    </svg>
  );
};

export {
  LockIcon,
  CloseIcon,
  CopyIcon,
  DeleteOrderIcon,
  DiscordIcon,
  EditOrderIcon,
  EditStakeIcon,
  KeyIcon,
  PauseIcon,
  PauseFilledIcon,
  PollIcon,
  ProposalIcon,
  SearchIcon,
  SendIcon,
  StakeIcon,
  SwapIcon,
  TransactionIcon,
  TwitterIcon,
  UnpauseIcon,
  UnstakeIcon,
  VoteIcon,
  WalletLogoIcon,
};
