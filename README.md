![logo](https://github.com/jesperkha/FigureSlice/.github/logo.png)

# **FigureSlice**

### **About**

Figure Slice is a web tool for clipping images with various shapes. The in-browser editor lets you create shapes that go over your image and act as a mask. When the image is processed you will only get the parts of the image that are seen through the shapes drawn. Changing the opacity of each shape is also possible.

![image of web editor](https://github.com/jesperkha/FigureSlice/.github/editor.jpg)

<br>

### **How to use**

The website has clear and consice instructions when it comes to actually using the editor.

To get the shape drawing p

<br>

### **How it works**

The editor is made with plain javascript and DOM manipulation. The shapes draw are just divs with a displaced background image relative to the source image. When done, the shapes positions and sizes are scaled to match the dimensions of the source image.

The code for the image editing process is found in the [img directory](https://github.com/jesperkha/FigureSlice/img). The shapes sent from the editor are drawn to a blank image buffer and that buffer is used as a mask when drawing the new image with the standard Go image/draw library.
