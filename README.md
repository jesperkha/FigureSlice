# **FigureSlice**

[Go to the website](https://figureslice.herokuapp.com)

<br>

## **About**

Figure Slice is a web tool for clipping images with various shapes. The in-browser editor lets you create shapes that go over your image and act as a mask. When the image is processed you will only get the parts of the image that are seen through the shapes drawn. Changing the opacity of each shape is also possible.

<br>

<img src="https://github.com/jesperkha/FigureSlice/blob/main/.github/editor.jpg?raw=true" alt="editor" width="70%"/>

<br>
<br>

## **How to use**

The website has clear and consice instructions when it comes to actually using the editor.
If something is unclear or you find a bug that is not listed below, please feel free to create an issue.

<br>

**Known bugs / issues**

- Circles will sometimes overlap the edge of the editor when drawing
- Currently only supports the `png` image format
- Opacity sliders value is not set to the shapes opacity when selected

<br>
<br>

## **How it works**

The editor is made with plain javascript and DOM manipulation. The shapes drawn are just divs with a displaced background image relative to the source image. When done, the shapes positions and sizes are scaled to match the dimensions of the source image.

The code for the image editing process is found in the [img directory](https://github.com/jesperkha/FigureSlice/tree/main/img). The shapes sent from the editor are drawn to a blank image buffer and that buffer is used as a mask when drawing the new image with the standard Go image/draw library.

<br>

<img src="https://github.com/jesperkha/FigureSlice/blob/main/.github/example.png?raw=true" alt="example" width="70%"/>
