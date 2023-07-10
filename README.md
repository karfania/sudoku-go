
<a name="readme-top"></a>

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <!-- <a href="https://github.com/karfania/sudoku-go">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

<h3 align="center">SUDOKU-GO</h3>

  <p align="center">
    A full-stack web application with playable Soduku. Leverages a Sudoku engine written in Go, a Gin API, and a React front-end.
    <br />
    <a href="https://github.com/karfania/sudoku-go"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <!-- <a href="https://github.com/karfania/sudoku-go">View Demo</a>
    ·
    <a href="https://github.com/karfania/sudoku-go/issues">Report Bug</a>
    ·
    <a href="https://github.com/karfania/sudoku-go/issues">Request Feature</a> -->
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

<!-- [![Product Name Screen Shot][product-screenshot]](https://example.com) -->
<!-- 

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->



### Built With

* [![React][React.js]][React-url]
* [![Typescript][Typescriptlang]][Typescript-url]
* [![Go][Golang]][Go-url]
* [![JQuery][JQuery.com]][JQuery-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started
This is a full-stack web application with plans to be deployed once development is complete. In the meantime, you can get a local copy running on your machine with the following steps:


### Prerequisites

* npm
  ```sh
  npm install npm@latest -g
  ```
* typescript
  ```sh
  npm install -g typescript
  ```
* golang
    ```
    https://go.dev/doc/install
    ```
### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/karfania/sudoku-go.git
   ```
2. Install NPM packages
   ```sh
   cd client
   ```
   ```sh
   npm install
   ```
3. Run local instance of the backend Sudoku engine & API
   ```sh
   cd server
   ```
   ```go
   go run main.go
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- ROADMAP -->
## Roadmap

- [ ] Sudoku Board Generator
  - [ ] Difficulty Settings
  - [ ] Timer
- [ ] Sudoku Move Validator
- [ ] Sudoku Board Solver

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

[Kourosh Arfania](https://karfania.github.io) - arfania@usc.edu

Project Link: [https://github.com/karfania/sudoku-go](https://github.com/karfania/sudoku-go)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/karfania/sudoku-go.svg?style=for-the-badge
[contributors-url]: https://github.com/karfania/sudoku-go/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/karfania/sudoku-go.svg?style=for-the-badge
[forks-url]: https://github.com/karfania/sudoku-go/network/members
[stars-shield]: https://img.shields.io/github/stars/karfania/sudoku-go.svg?style=for-the-badge
[stars-url]: https://github.com/karfania/sudoku-go/stargazers
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/karfania/
[product-screenshot]: images/screenshot.png
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[JQuery.com]: https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white
[JQuery-url]: https://jquery.com 
[Golang]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/
[Typescriptlang]: https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white
[Typescript-url]: https://www.typescriptlang.org
