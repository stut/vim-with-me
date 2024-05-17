package encoder

const (
    NONE = iota
    XOR_RLE
    XOR_BIT_DIFF // Not implement, but i am horned up for it
    HUFFMAN
    XOR_HUFFMAN
    HUFFMAN_QUADTREE // Maybe implement?
    XOR_RLE_QUADTREE // Maybe implement?
)
