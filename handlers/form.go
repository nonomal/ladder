package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func Form(c *fiber.Ctx) error {
	if os.Getenv("DISABLE_FORM") == "true" {
		c.Set("Content-Type", "text/html")
		c.SendStatus(fiber.StatusNotFound)
		return c.SendString("Form Disabled")
	} else {
		c.Set("Content-Type", "text/html")
		return c.SendString(html)
	}
}

const html = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Ladder</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css">
</head>
    <style>
        body {
            background-color: #ffffff;
        }

        header h1 {
            text-transform: uppercase;
            font-size: 70px;
            font-weight: 600;
            color: #fdfdfe;
            text-shadow: 0px 0px 5px #7AA7D1, 0px 0px 10px #7AA7D1, 0px 0px 10px #7AA7D1,
                0px 0px 20px #7AA7D1;
        }
        .logo-title {
            font-family: 'Arial', sans-serif;
            font-size: 2rem;
            color: #fff;
            margin-bottom: 20px;
        }
        .logo {
            text-align: center;
        }
        .github-corner {
            animation: octocat-wave 560ms ease-in-out;
            position:absolute;
            top:0;
            right:0;
        }
    </style>
</head>
<body>
    <a href="https://github.com/kubero-dev/ladder">
        <div class="github-corner" aria-label="View source on GitHub">
            <svg
                xmlns:svg="http://www.w3.org/2000/svg"
                xmlns="http://www.w3.org/2000/svg"
                xmlns:xlink="http://www.w3.org/1999/xlink"
                version="1.1"
                width="146"
                height="146"
                id="svg2">
                <defs
                    id="defs8">
                    <filter
                    height="1.096"
                    y="-0.048"
                    width="1.096"
                    x="-0.048"
                    style="color-interpolation-filters:sRGB"
                    id="filter6">
                    <feGaussianBlur
                        stdDeviation="3"
                        id="feGaussianBlur4" />
                    </filter>
                </defs>
                <path
                    d="M 152,140 6,-6 H 48 L 152,98 Z"
                    style="opacity:0.8;filter:url(#filter6)"
                    id="path10" />
                <path
                    d="M 146,134 12,0 h 42 l 92,92 z"
                    style="fill:#007200"
                    id="path12" />
                <g
                    aria-label="Fork me on GitHub"
                    transform="rotate(45)"
                    style="font-family:Collegiate;fill:#ffffff"
                    id="g42">
                    <path
                    d="m 53.643,-19.486 c 0,0.688 -0.016,1.2 -0.064,1.504 h 2.08 c -0.048,-0.32 -0.064,-0.8 -0.064,-1.424 v -3.344 h 1.76 c 0.416,0 0.736,0.016 0.944,0.048 v -1.76 c -0.24,0.032 -0.592,0.048 -1.088,0.048 h -1.616 v -2.496 h 1.936 c 0.56,0 0.944,0.016 1.184,0.048 v -1.792 h -5.136 c 0.048,0.272 0.064,0.784 0.064,1.504 z"
                    id="path14" />
                    <path
                    d="m 62.424,-17.87 c 1.008,0 1.776,-0.368 2.272,-1.088 0.432,-0.624 0.656,-1.472 0.656,-2.544 0,-2.416 -0.976,-3.616 -2.928,-3.616 -1.968,0 -2.96,1.2 -2.96,3.616 0,1.072 0.224,1.936 0.656,2.56 0.512,0.72 1.28,1.072 2.304,1.072 z m -0.016,-5.68 c 0.496,0 0.816,0.24 0.976,0.704 0.096,0.272 0.144,0.72 0.144,1.344 0,0.64 -0.048,1.088 -0.144,1.36 -0.16,0.464 -0.48,0.688 -0.976,0.688 -0.496,0 -0.816,-0.24 -0.976,-0.704 -0.096,-0.272 -0.144,-0.72 -0.144,-1.344 0,-0.624 0.048,-1.072 0.144,-1.344 0.16,-0.464 0.48,-0.704 0.976,-0.704 z"
                    id="path16" />
                    <path
                    d="m 68.293,-17.982 c -0.032,-0.24 -0.048,-0.64 -0.048,-1.184 v -3.888 c 0.352,-0.304 0.672,-0.464 0.976,-0.464 0.224,0 0.48,0.096 0.752,0.288 v -1.808 c -0.224,-0.08 -0.432,-0.128 -0.624,-0.128 -0.448,0 -0.832,0.192 -1.152,0.56 v -0.48 h -1.744 c 0.032,0.192 0.048,0.544 0.048,1.04 v 4.976 c 0,0.512 -0.016,0.88 -0.048,1.088 z"
                    id="path18" />
                    <path
                    d="m 72.857,-17.982 c -0.032,-0.24 -0.048,-0.64 -0.048,-1.184 v -2.448 l 1.472,2.816 c 0.208,0.384 0.32,0.656 0.368,0.816 h 1.872 l -2.352,-4.416 2.144,-2.752 h -2.064 c -0.08,0.176 -0.192,0.352 -0.352,0.56 l -1.088,1.44 v -4.496 c 0,-0.464 0.016,-0.8 0.048,-1.008 h -1.824 c 0.032,0.192 0.048,0.544 0.048,1.04 v 8.544 c 0,0.512 -0.016,0.88 -0.048,1.088 z"
                    id="path20" />
                    <path
                    d="m 85.08,-24.478 c -0.384,-0.432 -0.896,-0.656 -1.52,-0.656 -0.416,0 -0.864,0.192 -1.328,0.56 v -0.512 l -1.76,-0.016 c 0.032,0.176 0.048,0.544 0.048,1.12 v 4.992 c 0,0.496 -0.016,0.832 -0.048,1.008 h 1.856 c 0,-0.064 -0.048,-0.64 -0.048,-1.008 v -3.984 c 0.304,-0.288 0.608,-0.432 0.928,-0.432 0.656,0 0.864,0.416 0.864,1.76 l -0.016,2.16 c 0,0.656 -0.032,1.168 -0.08,1.504 h 1.92 c -0.048,-0.256 -0.064,-0.752 -0.064,-1.472 v -2.192 c 0,-0.56 -0.048,-1.056 -0.144,-1.504 0.208,-0.176 0.544,-0.256 0.976,-0.256 0.64,0 0.96,0.592 0.96,1.76 v 2.16 c 0,0.656 -0.032,1.168 -0.08,1.504 h 1.904 c -0.048,-0.256 -0.064,-0.752 -0.064,-1.472 v -2.192 c 0,-0.96 -0.176,-1.744 -0.512,-2.368 -0.432,-0.752 -1.056,-1.12 -1.888,-1.12 -0.736,0 -1.376,0.224 -1.904,0.656 z"
                    id="path22" />
                    <path
                    d="m 95.905,-20.99 c 0.032,-0.304 0.048,-0.624 0.048,-0.992 0,-0.944 -0.224,-1.696 -0.656,-2.256 -0.464,-0.592 -1.136,-0.896 -2.016,-0.896 -0.896,0 -1.6,0.368 -2.112,1.088 -0.464,0.656 -0.688,1.456 -0.688,2.432 0,1.136 0.272,2.048 0.832,2.704 0.576,0.72 1.392,1.072 2.448,1.072 0.496,0 1.056,-0.128 1.712,-0.368 v -1.712 c -0.464,0.304 -1.008,0.464 -1.6,0.464 -0.944,0 -1.44,-0.512 -1.52,-1.536 z m -2.576,-2.672 c 0.64,0 0.96,0.4 0.976,1.216 h -1.968 c 0.048,-0.816 0.368,-1.216 0.992,-1.216 z"
                    id="path24" />
                    <use
                    xlink:href="#path16"
                    transform="translate(40.438)"
                    id="use26" />
                    <path
                    d="m 110.187,-25.15 c -0.496,0 -0.992,0.208 -1.472,0.64 v -0.576 h -1.76 c 0.032,0.176 0.048,0.56 0.048,1.184 v 4.912 c 0,0.496 -0.016,0.832 -0.048,1.008 h 1.856 c 0,-0.064 -0.048,-0.64 -0.048,-1.008 v -3.936 c 0.368,-0.352 0.736,-0.528 1.088,-0.528 0.784,0 1.168,0.608 1.152,1.808 l -0.016,2.16 c -0.016,0.752 -0.032,1.264 -0.064,1.504 h 1.92 c -0.048,-0.256 -0.064,-0.752 -0.064,-1.472 v -2.192 c 0,-0.944 -0.192,-1.744 -0.592,-2.384 -0.464,-0.752 -1.136,-1.136 -2,-1.12 z"
                    id="path28" />
                    <path
                    d="m 123.877,-17.982 c 0.144,0.016 0.256,0.016 0.336,0 0,-0.192 -0.064,-0.768 -0.064,-1.36 v -1.856 c 0,-0.56 0.016,-1.056 0.064,-1.52 h -1.952 c 0.032,0.704 0.048,1.12 0.032,1.248 0,1.28 -0.512,1.92 -1.552,1.92 -1.264,0 -1.904,-1.264 -1.904,-3.776 0,-2.48 0.784,-3.728 2.352,-3.728 0.752,0 1.472,0.288 2.16,0.88 v -1.824 c -0.608,-0.528 -1.328,-0.8 -2.16,-0.8 -2.896,0 -4.416,1.904 -4.416,5.424 0,3.696 1.328,5.552 3.968,5.552 0.592,0 1.12,-0.128 1.584,-0.368 0.368,-0.208 0.624,-0.432 0.768,-0.688 z"
                    id="path30" />
                    <path
                    d="m 126.49,-26.334 c 0.592,0 1.104,-0.544 1.104,-1.184 0,-0.656 -0.512,-1.2 -1.104,-1.2 -0.624,0 -1.12,0.544 -1.12,1.2 0,0.64 0.496,1.184 1.12,1.184 z m 0.896,8.352 c -0.016,-0.24 -0.032,-0.64 -0.032,-1.184 v -4.912 c 0,-0.464 0.016,-0.8 0.032,-1.008 h -1.808 c 0.016,0.192 0.032,0.544 0.032,1.04 v 4.976 c 0,0.512 -0.016,0.88 -0.032,1.088 z"
                    id="path32" />
                    <path
                    d="m 130.783,-25.742 c 0,-0.256 0.016,-0.48 0.048,-0.688 h -1.856 c 0.032,0.176 0.048,0.416 0.048,0.72 v 0.624 h -0.784 v 1.552 c 0.224,-0.032 0.4,-0.048 0.544,-0.048 l 0.24,0.016 v 0.032 h -0.016 v 2.864 c 0,0.896 0.112,1.552 0.336,1.968 0.304,0.56 0.832,0.832 1.616,0.832 0.56,0 1.024,-0.112 1.424,-0.32 v -1.6 c -0.272,0.176 -0.56,0.272 -0.896,0.272 -0.464,0 -0.704,-0.352 -0.704,-1.072 v -2.976 h 0.688 c 0.256,0 0.592,0.032 0.704,0.032 v -1.552 h -1.392 z"
                    id="path34" />
                    <path
                    d="m 140.259,-27.678 c 0,-0.416 0.016,-0.736 0.064,-0.976 h -2.096 c 0.048,0.24 0.064,0.688 0.064,1.344 v 2.8 h -2.912 v -3.024 c 0,-0.48 0.016,-0.848 0.064,-1.12 h -2.08 c 0.048,0.256 0.064,0.624 0.064,1.12 v 8.432 c 0,0.496 -0.016,0.864 -0.064,1.12 h 2.08 c -0.048,-0.24 -0.064,-0.656 -0.064,-1.232 v -3.552 h 2.912 v 3.568 c 0,0.528 -0.016,0.944 -0.064,1.216 h 2.096 c -0.048,-0.24 -0.064,-0.624 -0.064,-1.12 v -3.664 h 0.528 v -1.744 h -0.528 z"
                    id="path36" />
                    <path
                    d="m 144.402,-17.918 c 0.56,0 1.072,-0.208 1.568,-0.64 v 0.576 h 1.744 c -0.016,-0.176 -0.032,-0.576 -0.032,-1.2 v -4.896 c 0,-0.496 0.016,-0.832 0.032,-1.008 h -1.856 c 0,0.048 0.064,0.64 0.064,1.008 v 3.936 c -0.368,0.352 -0.704,0.528 -1.008,0.528 -0.432,0 -0.72,-0.16 -0.88,-0.496 -0.144,-0.272 -0.208,-0.704 -0.208,-1.312 l 0.016,-2.16 c 0.016,-0.768 0.032,-1.264 0.064,-1.504 h -1.92 c 0.048,0.256 0.064,0.752 0.064,1.472 v 2.192 c 0,0.976 0.16,1.776 0.48,2.384 0.4,0.752 1.024,1.12 1.872,1.12 z"
                    id="path38" />
                    <path
                    d="m 152.31,-17.934 c 0.848,0 1.536,-0.416 2.048,-1.232 0.432,-0.704 0.64,-1.536 0.64,-2.48 0,-0.928 -0.208,-1.712 -0.608,-2.368 -0.48,-0.752 -1.136,-1.12 -1.984,-1.12 -0.464,0 -0.944,0.16 -1.44,0.464 v -2.784 c 0,-0.608 0.016,-1.008 0.032,-1.2 h -1.824 c 0.032,0.176 0.048,0.576 0.048,1.2 v 8.464 c 0,0.496 -0.016,0.832 -0.048,1.008 h 1.696 v -0.576 c 0.384,0.416 0.864,0.624 1.44,0.624 z m -0.24,-5.52 c 0.736,0 1.104,0.608 1.104,1.808 0,0.496 -0.08,0.928 -0.256,1.296 -0.208,0.464 -0.528,0.688 -0.944,0.688 -0.336,0 -0.672,-0.16 -1.008,-0.496 v -2.768 c 0.384,-0.352 0.752,-0.528 1.104,-0.528 z"
                    id="path40" />
                </g>
                <path
                    d="m 52,0 94,94 M 14,0 146,132"
                    style="fill:none;stroke:#ffffff;stroke-dasharray:2, 1;stroke-opacity:0.95"
                    id="path44" />
                </svg>
        </div>
    </a>
    <div class="container">
		<div class="logo">
			<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="250" height="250" viewBox="0 0 512 512">
				<path fill="#7AA7D1" d="M262.074 485.246C254.809 485.265 247.407 485.534 240.165 484.99L226.178 483.306C119.737 468.826 34.1354 383.43 25.3176 274.714C24.3655 262.975 23.5876 253.161 24.3295 241.148C31.4284 126.212 123.985 31.919 238.633 24.1259L250.022 23.8366C258.02 23.8001 266.212 23.491 274.183 24.1306C320.519 27.8489 366.348 45.9743 402.232 75.4548L416.996 88.2751C444.342 114.373 464.257 146.819 475.911 182.72L480.415 197.211C486.174 219.054 488.67 242.773 487.436 265.259L486.416 275.75C478.783 352.041 436.405 418.1 369.36 455.394L355.463 462.875C326.247 477.031 294.517 484.631 262.074 485.246ZM253.547 72.4475C161.905 73.0454 83.5901 144.289 73.0095 234.5C69.9101 260.926 74.7763 292.594 83.9003 317.156C104.53 372.691 153.9 416.616 211.281 430.903C226.663 434.733 242.223 436.307 258.044 436.227C353.394 435.507 430.296 361.835 438.445 267.978C439.794 252.442 438.591 236.759 435.59 221.5C419.554 139.955 353.067 79.4187 269.856 72.7052C264.479 72.2714 258.981 72.423 253.586 72.4127L253.547 72.4475Z"/>
				<path fill="#7AA7D1" d="M153.196 310.121L133.153 285.021C140.83 283.798 148.978 285.092 156.741 284.353L156.637 277.725L124.406 278.002C123.298 277.325 122.856 276.187 122.058 275.193L116.089 267.862C110.469 260.975 103.827 254.843 98.6026 247.669C103.918 246.839 105.248 246.537 111.14 246.523L129.093 246.327C130.152 238.785 128.62 240.843 122.138 240.758C111.929 240.623 110.659 242.014 105.004 234.661L97.9953 225.654C94.8172 221.729 91.2219 218.104 88.2631 214.005C84.1351 208.286 90.1658 209.504 94.601 209.489L236.752 209.545C257.761 209.569 268.184 211.009 285.766 221.678L285.835 206.051C285.837 197.542 286.201 189.141 284.549 180.748C280.22 158.757 260.541 143.877 240.897 135.739C238.055 134.561 232.259 133.654 235.575 129.851C244.784 119.288 263.68 111.99 277.085 111.105C288.697 109.828 301.096 113.537 311.75 117.703C360.649 136.827 393.225 183.042 398.561 234.866C402.204 270.253 391.733 308.356 367.999 335.1C332.832 374.727 269.877 384.883 223.294 360.397C206.156 351.388 183.673 333.299 175.08 316.6C173.511 313.551 174.005 313.555 170.443 313.52L160.641 313.449C158.957 313.435 156.263 314.031 155.122 312.487L153.196 310.121Z"/>
			</svg>
		</div>
        <header>
            <h1 class="center-align logo-title">ladddddddder</h1>
        </header>
        <form id="inputForm" class="col s12" method="get">
            <div class="row">
                <div class="input-field col s12">
                    <input type="text" id="inputField" name="inputField" class="validate" required>
                    <label for="inputField">URL</label>
                </div>
				<!--
                <div class="input-field col s2">
                    <button class="btn waves-effect waves-light" type="submit" name="action">Submit
                        <i class="material-icons right">go</i>
                    </button>
                </div>
				-->
            </div>
        </form>
    </div>

    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>
    <script>
        document.addEventListener('DOMContentLoaded', function() {
            M.AutoInit();
        });
        document.getElementById('inputForm').addEventListener('submit', function (e) {
            e.preventDefault();
            let url = document.getElementById('inputField').value;
            if (url.indexOf('http') === -1) {
                url = 'https://' + url;
            }
            window.location.href = '/' + url;
            return false;
        });
    </script>
</body>
</html>
`
