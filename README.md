[![MIT License][license-shield]][license-url]


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
        <li><a href="#prerequisite-steps">Prerequisite Steps</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li>
      <a href="#usage">Usage</a>
      <ul>
	<li><a href="#scheduling-data-collection">Scheduling Data Collection</a></li>
	<li><a href="#exporting-collected-data">Exporting Collected Data</a></li>
	<li><a href="#generating-heatmaps">Generating Heatmaps</a></li>
      </ul>
    </li>
    <li><a href="#contributing">Contributing</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

When designing new buildings, architects must consider many factors when designing the interior workspaces and designs of the various rooms. During this time, they will often examine pre existing buildings to manually create foot traffic maps, which they can then incorporate into their own designs. The same is expected of Architecture students, who are required to examine academic buildings around campus to create their own foot maps. However, between other classes and a limited budget, architecture students do not have the same resources that professionals would have. This is where the entity detection system comes in.

Tabula Plena will be an entity recognizing application with an intended use of generating foot traffic maps for architecture students. This will come in especially handy for students (and architects in general) as it will greatly contribute to the research process that architects will undertake when designing new buildings. Without software present, architects would previously have to rely on their own personal data collection to determine foot traffic. This can be a grueling experience which can lead to inaccurate results, as it requires extreme concentration over a long period of time. Our project aims to automate this process to reduce human error and allow data collection to be happening at all times of the day.


### Built With

* [GoCV](https://gocv.io/)


<!-- GETTING STARTED -->
## Getting Started

Everything you need to get your machine set up to use our software!

### Prerequisite Steps
1. Install Go (Golang): [Installation Guide](https://go.dev/doc/install)
2. Install the latest version of OpenCV: [Installation Guide](https://docs.opencv.org/4.x/df/d65/tutorial_table_of_content_introduction.html)
3. Install GoCV: [Installation Guide](https://gocv.io/getting-started/)

### Installation
1. Download or clone our [repository](https://github.com/pavanramadass/-Entity-Traffic-Map).
    - Instructions for cloning a GitHub repository can be found [here](https://docs.github.com/en/repositories/creating-and-managing-repositories/cloning-a-repository).
2. Use your operating system's command line to navigate to the repository's top level folder in your local file storage.
    - After opening the command line, run the command `cd $path` where `$path` is the file path leading to your copy of our repository.
3. Execute `go run main.go` in the command line to start the program.

<!-- USAGE -->
## Usage

### Scheduling Data Collection

### Exporting Collected Data

### Generating Heatmaps

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the project.
2. Create a new branch for your feature. (`git checkout -b feature/AmazingFeature`)
    - Make sure to follow the [Style Guide](https://go.dev/doc/effective_go).
3. Commit your changes. (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch. (`git push origin feature/AmazingFeature`)
5. Open a pull request.

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
