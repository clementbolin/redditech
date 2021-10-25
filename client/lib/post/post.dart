import 'package:flutter/material.dart';

class PostComponent extends StatefulWidget {
  const PostComponent({
    Key? key
  }): super(key:key);

  @override
  State<PostComponent> createState() => _PostComponent();
}

class _PostComponent extends State<PostComponent> {
  @override
  Widget build(BuildContext context) {
    return const Text("Hello Post");
  }
}