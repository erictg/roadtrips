import 'package:flutter/material.dart';
import 'package:roadtrips/screens/search.dart';

class HomeScreen extends StatelessWidget {
  static (String, Widget Function(BuildContext)) route =
      ('/', (context) => const HomeScreen());

  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Destinations')),
      body: Center(
        child:
            Column(mainAxisAlignment: MainAxisAlignment.spaceAround, children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.start,
            children: [
              const Expanded(
                flex: 1,
                child: Text(
                  'Where do you want to go to dinner tonight?',
                  style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
                ),
              ),
              Expanded(flex: 1, child: Container()),
            ],
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.end,
            children: [
              Expanded(flex: 1, child: Container()),
              const Expanded(
                flex: 1,
                child: Text(
                  'I don\'t know. Where do you want to go?',
                  style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
                ),
              ),
            ],
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.start,
            children: [
              const Expanded(
                flex: 1,
                child: Text(
                  'I don\'t know. How about one of these...',
                  style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
                ),
              ),
              Expanded(flex: 1, child: Container())
            ],
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.start,
            children: [
              Expanded(
                flex: 1,
                child: TextButton(
                    onPressed: () => Navigator.pushNamed(
                        context, '/destination/search',
                        arguments:
                            const SearchScreenArgs('/destination/suggestions')),
                    child: const Text(
                      'Suggestions',
                      style:
                          TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
                    )),
              ),
              Expanded(flex: 1, child: Container())
            ],
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.end,
            children: [
              Expanded(flex: 1, child: Container()),
              const Expanded(
                flex: 1,
                child: Text(
                  'Mmmm... what about something...',
                  style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
                ),
              ),
            ],
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.end,
            children: [
              Expanded(flex: 1, child: Container()),
              Expanded(
                  flex: 1,
                  child: TextButton(
                      onPressed: () => Navigator.pushNamed(
                          context, '/destination/search',
                          arguments:
                              const SearchScreenArgs('/destination/random')),
                      child: const Text(
                        'Random',
                        style: TextStyle(
                            fontSize: 20, fontWeight: FontWeight.bold),
                      ))),
            ],
          )
        ]),
      ),
    );
  }
}
