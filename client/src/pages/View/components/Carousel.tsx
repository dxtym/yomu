import React from "react";
import Pager from "./Pager";
import Loading from "@/components/common/Loading";

import { IChapter } from "@/types/chapter"
import { FC, ReactElement, useRef, useState } from "react";
import { Container, Box, Image } from "@chakra-ui/react";

const Carousel: FC<IChapter> = ({ page_urls }): ReactElement => {
  const [page, setPage] = useState<number>(0);
  const [zoom, setZoom] = useState<number>(1);
  const [pos, setPos] = useState({x: 0, y: 0});
  const imageRef = useRef<HTMLImageElement>(null);
  const containerRef = useRef<HTMLDivElement>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [dragging, setDragging] = useState<boolean>(false);

  const showNext = () => {
    if (page + 1 < page_urls.length) {
      setLoading(true);
      setPage(page + 1);
    }
  }

  const showPrev = () => {
    if (page - 1 >= 0) {
      setLoading(true);
      setPage(page - 1);
    }
  }

  const handleWheel = (e: React.WheelEvent<HTMLDivElement>) => {
    if (imageRef.current) {
      const newZoom = zoom + (e.deltaY > 0 ? -0.1 : 0.1);
      setZoom(Math.max(1, Math.min(newZoom, 3)));
    }
  }

  const handleMouseDown = (e: React.MouseEvent<HTMLDivElement>) => {
    setDragging(true);
    const initPos = { x: e.clientX, y: e.clientY };
    const initDragPos = { ...pos };

    const handleMouseMove = (e: MouseEvent) => {
      if (dragging) {
        const deltaX = e.clientX - initPos.x;
        const deltaY = e.clientY - initPos.y;
        setPos({
          x: initDragPos.x + deltaX,
          y: initDragPos.y + deltaY,
        });
      }
    };

    const handleMouseUp = () => {
      setDragging(false);
      document.removeEventListener('mousemove', handleMouseMove);
      document.removeEventListener('mouseup', handleMouseUp);
    };

    document.addEventListener('mousemove', handleMouseMove);
    document.addEventListener('mouseup', handleMouseUp);
  };

  return (
    <Container
      display={"flex"}
      justifyContent={"center"}
      flexDirection={"column"}
      height={"100vh"}
    >
      <Box
        py={"100%"}
        display={"flex"} 
        justifyContent={"center"} 
        alignItems={"center"}
        ref={containerRef}
        overflow={"hidden"}
        onWheel={handleWheel}
        onMouseDown={handleMouseDown}
      >
        {loading && <Loading />}
        <Image 
          ref={imageRef}
          src={page_urls[page]} 
          display={loading ? "none" : "block"} 
          onLoad={() => setLoading(false)}
          objectFit={"contain"}
          maxW={"100%"}
          height={"auto"}
          style={{
            transform: `scale(${zoom}) translate(${pos.x}px, ${pos.y}px)`,
            translate: `transform 0.2s ease`,
            cursor: dragging ? 'grabbing' : 'grab'
          }}
        />
      </Box>
      <Pager showNext={showNext} showPrev={showPrev} />
    </ Container>
  );
}

export default Carousel;