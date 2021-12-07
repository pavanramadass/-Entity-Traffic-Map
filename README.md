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
      <a href="#features">Features</a>
      <ul>
	<li><a href="#simple-intuitive-ui">Simple, Intuitive UI</a></li>
	<li><a href="#automated-schedulable-data-collection">Automated, Schedulable Data Collection</a></li>
	<li><a href="#raw-data-exportation-for-custom-analysis">Raw Data Exportation for Custom Analysis</a></li>
	<li><a href="#heatmap-generation">Heatmap Generation</a></li>
      </ul>
    </li>
    <li>
      <a href="#release-notes">Release Notes</a>
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

<!-- FEATURES -->
## Features

### Simple, Intuitive UI
Tabula Plena's UI is designed to allow a user with basic computer knowledge to easily utilize all our provided features without requiring an in-depth tutorial. The UI features simple, self-explanatory options and visual indicators of successful use.

### Automated, Schedulable Data Collection
Tabula Plena's primary feature is the automation of the collection of foot traffic data in buildings. Data collection can be scheduled for a set time period and the software will handle the rest, using computer vision to collect foot traffic data from a top-down video feed. Scheduled collection times can be modified, and active collection can be cancelled; the collected data is stored by the software when a scheduled collection ends, or when an active one is cancelled. 

### Raw Data Exportation for Custom Analysis
Tabula Plena stores collected foot traffic data by associating a Unix UTC timestamp with X and Y coordinates. The timestamp is taken when an entity is detected by the software, and the X and Y coordinates correspond to the pixel location of the detected entity's centroid in the camera's FOV. Provided metadata includes start and end times, as well as a base image taken from the camera feed. Tabula Plena allows users to download collected data in .json format for their own use and custom analysis.

> ***PRIVACY STATEMENT:*** Outside of the single base image of the camera feed, no actual video images are saved. Only the timestamp of detection and X and Y pixel coordinates are saved.

### Heatmap Generation 
One possible use of the data provided by Tabula Plena is the generation of foot traffic heatmaps. Tabula Plena offers a foot traffic heatmap generation tool for in-app analysis of collected data. Most data points are displayed as green dots, with higher traffic areas shown via warmer tones approaching red. The heatmap can be overlayed on the base image included with the metadata to put the data in context.

<!-- RELEASE NOTES -->
## Release Notes


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
